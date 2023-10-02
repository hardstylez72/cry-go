package repository

import (
	"context"

	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/server/repository/pg"
	"github.com/pkg/errors"
)

type ProfileRelation struct {
	P1Id      string `db:"p1_id"`
	P1SubType string `db:"p1_sub_type"`
	P1Type    string `db:"p1_type"`

	P2Id      string `db:"p2_id"`
	P2SubType string `db:"p2_sub_type"`
	P2Type    string `db:"p2_type"`

	UserId string `db:"user_id"`
}

type ProfileRel struct {
	Id      string `db:"id"`
	SubType string `db:"sub_type"`
	Type    string `db:"type"`
	Num     int64  `db:"num"`
}

type GetProfileRelationsRes struct {
	ProfileRel
	Related *ProfileRel
}

type GetProfileRelationsReq struct {
	P1Type, P2Type       v1.ProfileType
	P1SubType, P2SubType v1.ProfileSubType
	UserId               string
}

func (r *pgRepository) GetProfileRelations(ctx context.Context, req *GetProfileRelationsReq) ([]GetProfileRelationsRes, error) {

	relationMap, err := r.profileRelations(ctx, req.P1Type, req.P2Type, req.P1SubType, req.P2SubType, req.UserId)
	if err != nil {
		return nil, err
	}

	profiles, err := r.profiles(ctx, req.P1Type, req.P1SubType, req.UserId)
	if err != nil {
		return nil, err
	}

	out := make([]GetProfileRelationsRes, 0)

	for _, p := range profiles {
		h := GetProfileRelationsRes{
			ProfileRel: ProfileRel{
				Id:      p.Id,
				SubType: p.SubType,
				Type:    p.Type,
				Num:     p.Num,
			},
			Related: nil,
		}
		v, ok := relationMap[p.Id]
		if ok {
			h.Related = &ProfileRel{
				Id:      v.P2Id,
				SubType: v.P2SubType,
				Type:    v.P2Type,
				Num:     v.Num,
			}
		}

		out = append(out, h)
	}

	return out, nil
}

func (r *pgRepository) profiles(ctx context.Context, p1Type v1.ProfileType, p1SubType v1.ProfileSubType, userId string) ([]ProfileRel, error) {
	q := `select id,num, type, sub_type from profiles where type = $1 and sub_type = $2 and user_id = $3 and deleted_at is null`
	out := make([]ProfileRel, 0)

	if err := r.conn.SelectContext(ctx, &out, q, p1Type.String(), p1SubType.String(), userId); err != nil {
		return nil, pg.PgError(err)
	}

	return out, nil
}

type ProfileRelationSpecial struct {
	ProfileRelation
	Num int64 `db:"num"`
}

func (r *pgRepository) profileRelations(ctx context.Context, p1Type, p2Type v1.ProfileType, p1SubType, p2SubType v1.ProfileSubType, userId string) (map[string]ProfileRelationSpecial, error) {
	q := `select pr.*, p.num from profile_relations pr 
               join profiles p on (pr.p2_id = p.id)
               where p1_type = $1 and p2_type = $2
			and p1_sub_type = $3 and p2_sub_type = $4
			and pr.user_id = $5 and p.deleted_at is null`
	out := make([]ProfileRelationSpecial, 0)

	if err := r.conn.SelectContext(ctx, &out, q, p1Type.String(), p2Type.String(), p1SubType.String(), p2SubType.String(), userId); err != nil {
		return nil, pg.PgError(err)
	}

	m := make(map[string]ProfileRelationSpecial)

	for _, t := range out {
		m[t.P1Id] = t
	}

	return m, nil
}

func (r *pgRepository) AddProfileRelation(ctx context.Context, req *ProfileRelation) error {
	q := `insert into  profile_relations (p1_id, p1_sub_type, p1_type, p2_id, p2_sub_type, p2_type, user_id)
    values (:p1_id, :p1_sub_type, :p1_type, :p2_id, :p2_sub_type, :p2_type, :user_id)`
	if _, err := r.conn.NamedExecContext(ctx, q, req); err != nil {
		return pg.PgError(err)
	}
	return nil
}
func (r *pgRepository) DeleteProfileRelation(ctx context.Context, req *ProfileRelation) error {
	q := `delete from  profile_relations 
    where p1_id = :p1_id and p1_sub_type = :p1_sub_type and p1_type = :p1_type
	and  p2_id = :p2_id and p2_sub_type = :p2_sub_type and p2_type = :p2_type and user_id = :user_id`
	if _, err := r.conn.NamedExecContext(ctx, q, req); err != nil {
		return pg.PgError(err)
	}
	return nil
}

type SearchProfilesNotRelated struct {
	PType    string `db:"p_type"`
	PSubType string `db:"p_sub_type"`

	NotPType    string `db:"not_p_type"`
	NotPSubType string `db:"not_p_sub_type"`

	UserId string `db:"user_id"`
}

func (r *pgRepository) SearchProfilesNotRelated(ctx context.Context, req *SearchProfilesNotRelated) ([]ProfileRel, error) {

	//todo: в обе стороны
	q := `
	select 
	    p.id, 
	    p.type, 
	    p.sub_type,
	    p.num 
	from profiles p
		where p.id not in (select ttt.id from (
		    	    select p2_id as id from profile_relations
	                 where p1_sub_type = :p_sub_type
	                 and p1_type = :p_type
	                 and p2_sub_type = :not_p_sub_type
	                 and p2_type = :not_p_type
	                 and user_id = :user_id
	    union 
	    	    select p2_id as id from profile_relations
	                 where p1_sub_type = :not_p_sub_type
	                 and p1_type = :not_p_type
	                 and p2_sub_type = :p_sub_type
	                 and p2_type = :p_type
	                 and user_id = :user_id
		) ttt)
		and p.type = :p_type 
		and p.sub_type = :p_sub_type
		and p.deleted_at is null
		and p.user_id = :user_id
		order by p.num asc
   `
	out := make([]ProfileRel, 0)

	rows, err := r.conn.NamedQueryContext(ctx, q, req)
	if err != nil {
		return nil, pg.PgError(err)
	}

	for rows.Next() {
		var item ProfileRel
		if err := rows.StructScan(&item); err != nil {
			return nil, err
		}
		out = append(out, item)
	}

	return out, nil
}

type GetRelatedProfileReq struct {
	ProfileId      string
	ProfileType    string
	ProfileSubType string

	WantProfileType    string
	WantProfileSubType string

	UserId string
}

func (r *pgRepository) GetRelatedProfile(ctx context.Context, req *GetRelatedProfileReq) (*string, error) {

	id, err := r.getRelatedProfile(ctx, req)
	if err != nil {
		if errors.Is(err, pg.EntityNotFound) {
			return r.getRelatedProfileReversed(ctx, req)
		}
		return nil, err
	}

	return id, nil
}
func (r *pgRepository) getRelatedProfileReversed(ctx context.Context, req *GetRelatedProfileReq) (*string, error) {
	q := `select p1_id from  profile_relations 
             where p2_id = $1 
                 and p2_sub_type = $2
                     and p2_type = $3
                         and p1_sub_type =$4
                             and p1_type = $5
                             	and user_id = $6`

	var id string
	if err := r.conn.GetContext(ctx, &id, q, req.ProfileId, req.ProfileSubType, req.ProfileType, req.WantProfileSubType, req.WantProfileType, req.UserId); err != nil {
		return nil, pg.PgError(err)
	}

	return &id, nil
}
func (r *pgRepository) getRelatedProfile(ctx context.Context, req *GetRelatedProfileReq) (*string, error) {
	q := `select p2_id from  profile_relations 
             where p1_id = $1 
                 and p1_sub_type = $2
                     and p1_type = $3
                         and p2_sub_type =$4
                             and p2_type = $5
                             	and user_id = $6`

	var id string
	if err := r.conn.GetContext(ctx, &id, q, req.ProfileId, req.ProfileSubType, req.ProfileType, req.WantProfileSubType, req.WantProfileType, req.UserId); err != nil {
		return nil, pg.PgError(err)
	}

	return &id, nil
}
