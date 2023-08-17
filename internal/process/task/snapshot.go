package task

import (
	"context"
	"time"

	"github.com/hardstylez72/cry/internal/log"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/snapshot"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type SnapshotVoteTask struct {
	cancel func()
}

func (t *SnapshotVoteTask) Stop() error {
	t.cancel()
	return nil
}

func (t *SnapshotVoteTask) Type() v1.TaskType {
	return v1.TaskType_SnapshotVote
}

func (t *SnapshotVoteTask) Run(ctx context.Context, a *Input) (*v1.ProcessTask, error) {

	taskContext, cancel := context.WithTimeout(ctx, taskTimeout)
	defer cancel()

	t.cancel = cancel

	task := a.Task
	l, ok := a.Task.Task.Task.(*v1.Task_SnapshotVoteTask)
	if !ok {
		return nil, errors.New("panic.a.Task.Task.Task.(*v1.Task_SnapshotVoteTask) call an ambulance!")
	}

	p := l.SnapshotVoteTask

	profile, err := a.Halper.Profile(ctx, a.ProfileId)
	if err != nil {
		return nil, err
	}

	s, err := profile.GetNetworkSettings(ctx, v1.Network_ARBITRUM)
	if err != nil {
		return nil, err
	}

	switch a.Task.Status {
	case v1.ProcessStatus_StatusDone, v1.ProcessStatus_StatusError:
		return a.Task, nil
	case v1.ProcessStatus_StatusRetry, v1.ProcessStatus_StatusReady, v1.ProcessStatus_StatusRunning:

		before := a.Task.Status
		after := v1.ProcessStatus_StatusRunning
		a.Task.Status = after
		if err := a.ProcessRepository.UpdateProcessTask(ctx, task, task.Id, a.ProcessId, a.ProfileId); err != nil {
			return nil, errors.Wrap(err, "UpdateProcessTask")
		}
		if err := a.ProcessRepository.RecordStatusChanged(ctx, task.Id, before, after); err != nil {
			log.Log.Warn("RecordStatusChanged", zap.Error(err))
		}
	}

	if p.Proposal == nil {
		proposals, err := a.Snapshot.ActiveProposals(taskContext, &snapshot.ActiveProposalsReq{
			ProviderRPC: s.BaseConfig().RPCEndpoint,
			Space:       p.Space,
			Pk:          profile.WalletPK,
		})
		if err != nil {
			return nil, errors.Wrap(err, "Snapshot.ActiveProposals")
		}

		p.Proposal = make(map[string]*v1.SnapshotVoteProposal)

		for _, proposal := range proposals {
			p.Proposal[proposal.Id] = &v1.SnapshotVoteProposal{
				Status: v1.ProcessStatus_StatusReady,
				Msg:    "",
				Link:   "https://snapshot.org/#/stgdao.eth/proposal/" + proposal.Id,
				Id:     proposal.Id,
				Type:   proposal.Type,
				Symbol: proposal.Symbol,
			}
		}
		if err := a.ProcessRepository.UpdateProcessTask(ctx, task, a.Task.Id, a.ProcessId, a.ProfileId); err != nil {
			return nil, errors.Wrap(err, "UpdateProcessTask")
		}
	}

	for proposalId, proposal := range p.Proposal {
		if proposal.Status == v1.ProcessStatus_StatusDone {
			continue
		}

		proposal.Status = v1.ProcessStatus_StatusRunning
		p.Proposal[proposalId] = proposal

		if err := a.ProcessRepository.UpdateProcessTask(ctx, task, a.Task.Id, a.ProcessId, a.ProfileId); err != nil {
			return nil, errors.Wrap(err, "UpdateProcessTask")
		}

		networks := []v1.Network{
			v1.Network_ARBITRUM,
		}

		for _, network := range networks {
			s, err := profile.GetNetworkSettings(ctx, network)
			if err != nil {
				return nil, err
			}

			VoteContext, VoteContextCancel := context.WithTimeout(ctx, time.Minute*2)
			err = a.Snapshot.Vote(VoteContext, &snapshot.VoteReq{
				Type:        proposal.Type,
				ProposalId:  proposalId,
				ProviderRPC: s.BaseConfig().RPCEndpoint,
				Space:       p.Space,
				Pk:          profile.WalletPK,
			})
			VoteContextCancel()

			if err != nil {
				e := err
				proposal.Status = v1.ProcessStatus_StatusError
				proposal.Msg = e.Error()
				p.Proposal[proposalId] = proposal
				if err := a.ProcessRepository.UpdateProcessTask(ctx, task, a.Task.Id, a.ProcessId, a.ProfileId); err != nil {
					return nil, errors.Wrap(err, "UpdateProcessTask")
				}
				return nil, e
			}

			proposal.Status = v1.ProcessStatus_StatusDone
			proposal.Msg = "ok"
			p.Proposal[proposalId] = proposal
			if err := a.ProcessRepository.UpdateProcessTask(ctx, task, a.Task.Id, a.ProcessId, a.ProfileId); err != nil {
				return nil, errors.Wrap(err, "UpdateProcessTask")
			}

			break
		}
	}
	task.Status = v1.ProcessStatus_StatusDone
	task.FinishedAt = timestamppb.Now()

	if err := a.ProcessRepository.UpdateProcessTask(ctx, task, a.Task.Id, a.ProcessId, a.ProfileId); err != nil {
		return nil, errors.Wrap(err, "UpdateProcessTask")
	}

	if err := a.ProcessRepository.RecordStatusChanged(ctx, task.Id, v1.ProcessStatus_StatusRunning, v1.ProcessStatus_StatusDone); err != nil {
		log.Log.Warn("RecordStatusChanged", zap.Error(err))
	}

	return task, nil
}
