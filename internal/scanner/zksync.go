package scanner

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/hardstylez72/cry/internal/defi/zksyncera"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
)

type zkSync struct {
	Cli *http.Client
}

func NewScannerZkSync(c *http.Client) *zkSync {
	return &zkSync{
		Cli: c,
	}
}

func (s *zkSync) Transfers(ctx context.Context, add string) ([]Transfer, error) {

	limit := 100
	page := 1

	packs := []TransferRes{}

	for {
		tmp, err := s.transferPack(ctx, add, limit, page)
		if err != nil {
			return nil, err
		}
		packs = append(packs, *tmp)

		if page >= tmp.Meta.TotalPages {
			break
		}
		page++
	}

	out := make([]Transfer, 0)

	for _, p := range packs {
		for _, el := range p.Items {

			am := big.NewInt(0)
			if el.Amount != nil {
				v, ok := big.NewInt(0).SetString(*el.Amount, 10)
				if ok {
					am = big.NewInt(0).Add(am, v)
				}
			}

			out = append(out, Transfer{
				From:   el.From,
				Token:  el.TokenAddress,
				To:     el.To,
				Amount: am,
			})
		}
	}

	return out, nil
}
func (s *zkSync) Transactions(ctx context.Context, add string) ([]Transaction, error) {

	limit := 100
	page := 1

	packs := []TransactionRes{}

	for {
		tmp, err := s.transactionPack(ctx, add, limit, page)
		if err != nil {
			return nil, err
		}
		packs = append(packs, *tmp)

		if page >= tmp.Meta.TotalPages {
			break
		}
		page++
	}

	out := make([]Transaction, 0)

	for _, p := range packs {
		for _, el := range p.Items {
			out = append(out, Transaction{
				To: el.To,
				TS: el.ReceivedAt.Unix(),
			})
		}
	}

	return out, nil
}

func (s *zkSync) transferPack(ctx context.Context, addr string, limit, page int) (*TransferRes, error) {
	url := strings.Join([]string{
		"https://block-explorer-api.mainnet.zksync.io/address/",
		addr,
		"/transfers?",
		"limit=", strconv.Itoa(limit),
		"&page=", strconv.Itoa(page),
	}, "")

	return Request[any, TransferRes](ctx, s.Cli, http.MethodGet, url, nil)
}

func (s *zkSync) transactionPack(ctx context.Context, addr string, limit, page int) (*TransactionRes, error) {

	url := strings.Join([]string{
		"https://block-explorer-api.mainnet.zksync.io/transactions?address=",
		addr,
		"&limit=", strconv.Itoa(limit),
		"&page=", strconv.Itoa(page),
	}, "")

	return Request[any, TransactionRes](ctx, s.Cli, http.MethodGet, url, nil)
}

func Request[Req any, Res any](ctx context.Context, cli *http.Client, method string, url string, req *Req) (*Res, error) {

	marshal, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	var reqBody io.Reader
	if method != http.MethodGet {
		reqBody = bytes.NewBuffer(marshal)
	} else {
		reqBody = nil
	}

	r, err := http.NewRequestWithContext(ctx, method, url, reqBody)
	if err != nil {
		return nil, err
	}
	r.Header.Set("Content-Type", "application/json")
	res, err := cli.Do(r)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(string(body))
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var ser Res
	if err := json.Unmarshal(body, &ser); err != nil {
		return nil, err
	}

	return &ser, nil
}

type TransactionRes struct {
	Items []struct {
		Hash             string    `json:"hash"`
		To               string    `json:"to"`
		From             string    `json:"from"`
		Data             string    `json:"data"`
		Value            string    `json:"value"`
		IsL1Originated   bool      `json:"isL1Originated"`
		Fee              string    `json:"fee"`
		Nonce            int       `json:"nonce"`
		BlockNumber      int       `json:"blockNumber"`
		L1BatchNumber    int       `json:"l1BatchNumber"`
		BlockHash        string    `json:"blockHash"`
		TransactionIndex int       `json:"transactionIndex"`
		ReceivedAt       time.Time `json:"receivedAt"`
		Status           string    `json:"status"`
		CommitTxHash     string    `json:"commitTxHash"`
		ExecuteTxHash    string    `json:"executeTxHash"`
		ProveTxHash      string    `json:"proveTxHash"`
		IsL1BatchSealed  bool      `json:"isL1BatchSealed"`
	} `json:"items"`
	Meta struct {
		TotalItems   int `json:"totalItems"`
		ItemCount    int `json:"itemCount"`
		ItemsPerPage int `json:"itemsPerPage"`
		TotalPages   int `json:"totalPages"`
		CurrentPage  int `json:"currentPage"`
	} `json:"meta"`
	Links struct {
		First    string `json:"first"`
		Previous string `json:"previous"`
		Next     string `json:"next"`
		Last     string `json:"last"`
	} `json:"links"`
}

type TransferRes struct {
	Items []struct {
		From            string    `json:"from"`
		To              string    `json:"to"`
		BlockNumber     int       `json:"blockNumber"`
		TransactionHash string    `json:"transactionHash"`
		Timestamp       time.Time `json:"timestamp"`
		Amount          *string   `json:"amount"`
		TokenAddress    string    `json:"tokenAddress"`
		Type            string    `json:"type"`
		Fields          *struct {
			TokenId string `json:"tokenId"`
		} `json:"fields"`
		Token *struct {
			L2Address string  `json:"l2Address"`
			L1Address *string `json:"l1Address"`
			Symbol    string  `json:"symbol"`
			Name      string  `json:"name"`
			Decimals  int     `json:"decimals"`
		} `json:"token"`
	} `json:"items"`
	Meta struct {
		TotalItems   int `json:"totalItems"`
		ItemCount    int `json:"itemCount"`
		ItemsPerPage int `json:"itemsPerPage"`
		TotalPages   int `json:"totalPages"`
		CurrentPage  int `json:"currentPage"`
	} `json:"meta"`
	Links struct {
		First    string `json:"first"`
		Previous string `json:"previous"`
		Next     string `json:"next"`
		Last     string `json:"last"`
	} `json:"links"`
}

func (s *zkSync) SpecMap() map[string]v1.TaskType {

	out := make(map[string]v1.TaskType)

	for _, spec := range zksyncera.SpecMap {
		out[spec.Addr.String()] = spec.TaskType
	}

	return out
}

func (s *zkSync) Token(a string) (*v1.Token, bool) {
	out := map[address]v1.Token{}
	for token, addr := range zksyncera.TokenAddress {
		out[addr.String()] = token
	}

	out["0x000000000000000000000000000000000000800A"] = v1.Token_ETH

	v, ok := out[a]
	if ok {
		return &v, true
	}
	return nil, false
}

func (s *zkSync) Service(a string) (*Service, bool) {

	syncSwap := Service{
		Name: "SyncSwap",
		Url:  "https://syncswap.xyz/",
	}

	Holdstation := Service{
		Name: "Holdstation",
		Url:  "https://holdstation.exchange/",
	}

	iZUMifinance := Service{
		Name: "iZUMi finance",
		Url:  "https://izumi.finance/home",
	}

	Maverick := Service{
		Name: "Maverick",
		Url:  "https://app.mav.xyz/",
	}

	Muteio := Service{
		Name: "Mute.io",
		Url:  "https://mute.io/",
	}

	OnchainTrade := Service{
		Name: "Onchain Trade",
		Url:  "https://onchain.trade/",
	}

	OrbiterFinance := Service{
		Name: "Orbiter Finance",
		Url:  "https://orbiter.finance/",
	}

	SpaceFi := Service{
		Name: "SpaceFi",
		Url:  "https://spacefi.io/",
	}

	Starmaker := Service{
		Name: "Starmaker",
		Url:  "https://starmaker.exchange/",
	}

	Velocore := &Service{
		Name: "Velocore",
		Url:  "https://velocore.xyz/",
	}

	zkSyncEraPortal := &Service{
		Name: "zkSync Era Portal",
		Url:  "https://portal.zksync.io/bridge",
	}

	zkSyncNameService := &Service{
		Name: "zkSync Name Service",
		Url:  "https://app.zkns.domains/",
	}

	Goal3 := &Service{
		Name: "Goal3",
		Url:  "https://beta.goal3.xyz/",
	}

	zkSyncID := &Service{
		Name: "zkSync ID",
		Url:  "https://www.zksyncid.xyz/",
	}

	ZkSwap := &Service{
		Name: "ZkSwap",
		Url:  "https://zkswap.finance/",
	}

	XYFinance := &Service{
		Name: "XY Finance",
		Url:  "https://xy.finance/",
	}

	Ezkalibur := &Service{
		Name: "Ezkalibur",
		Url:  "https://dapp.ezkalibur.com",
	}

	PancakeSwap := &Service{
		Name: "PancakeSwap",
		Url:  "https://pancakeswap.finance/?chain=zkSync",
	}

	BigInt := &Service{
		Name: "BigInt",
		Url:  "https://bigint.co/",
	}

	RollupFinance := &Service{
		Name: "Rollup Finance",
		Url:  "https://rollup.finance/",
	}

	Basilisk := &Service{
		Name: "Basilisk",
		Url:  "https://app.basilisk.org/",
	}

	Odos := &Service{
		Name: "Odos",
		Url:  "https://app.odos.xyz/",
	}

	Tevaera := &Service{
		Name: "Tevaera",
		Url:  "https://tevaera.com/",
	}

	DraculaFi := &Service{
		Name: "DraculaFi",
		Url:  "https://draculafi.xyz/",
	}

	m := map[address]*Service{
		"0x80115c708e12edd42e504c1cd52aea96c547c05c": &syncSwap,
		"0x176b23f1ddfeedd10fc250b8f769362492ef810b": &syncSwap,
		"0x4e7d2e3f40998daeb59a316148bfbf8efd1f7f3c": &syncSwap,
		"0xae092fcec5fd04836b12e87da0d7ed3a707b38b0": &syncSwap,
		"0x2da10a1e27bf85cedd8ffb1abbe97e53391c0295": &syncSwap,
		"0x621425a1Ef6abE91058E9712575dcc4258F8d091": &syncSwap,

		"0x7b4872e2096ec9410b6b8c8b7d039589e6ee8022": &Holdstation,
		"0xaf08a9d918f16332f22cf8dc9abe9d9e14ddcbc2": &Holdstation,
		"0x51956481ced6f65458a25207e725cbd2e33a02cb": &Holdstation,
		"0x19b1ae9fd718107da3bed1f39ed9b3a5adbe89c0": &Holdstation,

		"0x943ac2310d9bc703d6ab5e5e76876e212100f894": &iZUMifinance,
		"0x83109541318bbd104c9c85328696107f5767dffb": &iZUMifinance,
		"0x483fde31bce3dcc168e23a870831b50ce2ccd1f1": &iZUMifinance,

		"0xfd54762d435a490405dda0fbc92b7168934e8525": &Maverick,
		"0x852639ee9dd090d30271832332501e87d287106c": &Maverick,
		"0x77ee88b1c9cce741ec35553730eb1f19cd45a379": &Maverick,
		"0x39e098a153ad69834a9dac32f0fca92066ad03f4": &Maverick,

		"0xdfaab828f5f515e104baaba4d8d554da9096f0e4": &Muteio,
		"0xb85feb6af3412d690dfda280b73eaed73a2315bc": &Muteio,
		"0x8b791913eb07c32779a16750e3868aa8495f5964": &Muteio,

		"0x84c18204c30da662562b7a2c79397c9e05f942f0": &OnchainTrade,
		"0xca806b267fc0c1c12edbf77a2e5bca5939c7d953": &OnchainTrade,
		"0xaa08a6d7b10724d03b8f4857d4fa14e7f92814e3": &OnchainTrade,
		"0x10c8044ae3f2b1c7decb439ff2dc1164d750c39d": &OnchainTrade,
		"0xe676b11421d68a28ba50920f2841183af93067a2": &OnchainTrade,
		"0x6219f06135b79705d34f5261852e9f6b98821e1f": &OnchainTrade,

		"0xe4edb277e41dc89ab076a1f049f4a3efa700bce8": &OrbiterFinance,
		"0x80c67432656d59144ceff962e8faf8926599bcf8": &OrbiterFinance,
		"0xee73323912a4e3772b74ed0ca1595a152b0ef282": &OrbiterFinance,
		"0x0a88BC5c32b684D467b43C06D9e0899EfEAF59Df": &OrbiterFinance,

		"0x0700fb51560cfc8f896b2c812499d17c5b0bf6a7": &SpaceFi,
		"0xe8826fc3ce6e74932144dbc2b369e7b52e88193a": &SpaceFi,
		"0x7cf85f98c0339559eab22deea1e018721e800add": &SpaceFi,
		"0xb376fceacd9fef24a342645cbf72a4c439ea0614": &SpaceFi,
		"0xacf5a67f2fcfeda3946ccb1ad9d16d2eb65c3c96": &SpaceFi,
		"0x4ad9ee1698b6d521ebb2883dd9fc3655c7553561": &SpaceFi,
		"0x00f093ff2bca9da894396336286c7c5bd2310ca5": &SpaceFi,
		"0x307baa142ba2bfa4a3059efb631899c992a193ee": &SpaceFi,
		"0x77d807b74d54b81a87a5769176bc7719f676c8ce": &SpaceFi,
		"0xbe7d1fd1f6748bbdefc4fbacafbb11c6fc506d1d": &SpaceFi,

		"0x1bdb8250eaf3c596441e6c3417c9d5195d6c85b9": &Starmaker,

		"0xcd52cbc975fbb802f82a1f92112b1250b5a997df": Velocore,
		"0xd999e16e68476bc749a28fc14a0c3b6d7073f50c": Velocore,

		"0x000000000000000000000000000000000000800a": zkSyncEraPortal,

		"0xae23b6e7f91ddebd3b70d74d20583e3e674bd94f": zkSyncNameService,

		"0xd2ca21df45479824f954a6e1759d436a57d63faf": Goal3,
		"0x1f090f91ee09768ca36dcd52640f4a5eae146555": Goal3,
		"0xc523df658dbec88dc03fb23a703bcbd7ffb5ea01": Goal3,
		"0x116a4a5ed4c7d5712e109d4188e17616d8e5784a": Goal3,
		"0x8d123a2a0a7c98555931ceda6917b865b7345164": Goal3,

		"0xa531ece441138d8ce52642ad497059f2a79fd96f": zkSyncID,

		"0x7642e38867860d4512fcce1116e2fb539c5cdd21": ZkSwap,
		"0xf100ff84b363af74e3fcdff554e3da3309159458": ZkSwap,
		"0x4eaaab540065b4e404d79fe1f0bf3a9c046f9151": ZkSwap,
		"0x6f89797fe9c97e57a1137fa70caa69f8abbfff50": ZkSwap,
		"0x18381c0f738146fb694de18d1106bde2be040fa4": ZkSwap,

		"0x936fef4245f281ed4f2ee303060a8656399dcc32": XYFinance,
		"0xada6a5c2ddf499b50efced4c489a90b707693ac5": XYFinance,
		"0xaa0b0654e79e17332d983e2351bd926ce336b9bd": XYFinance,
		"0x935283a00fbf8e40fd2f8c432a488f6addc8db67": XYFinance,
		"0xdf4700daccbfb097204576cbbb4728d4f190fdc1": XYFinance,
		"0x75167284361c8d61be7e4402f4953e2b112233cb": XYFinance,
		"0xc81dc052661785c5a304ebb2c4bceadcfb79a675": XYFinance,
		"0x9e52c2bd0177404dca3e122992a7fc0bcb15fd95": XYFinance,
		"0x83529b846762d1146c251548a3d8e453c1d09424": XYFinance,
		"0xbcdde9ca12f1141943e98249dae5b18a3d885e6b": XYFinance,
		"0x48ca44863861fd0687b3a21211bf90cb5924704e": XYFinance,
		"0xe4e156167cc9c7ac4abd8d39d203a5495f775547": XYFinance,
		"0x30e63157bd0ba74c814b786f6ea2ed9549507b46": XYFinance,

		"0x17a48ce80a7f7f3f774390179b1a404ce9c9a77e": Ezkalibur,
		"0x15c664a62086c06d43e75bb3fdded93008b8ce63": Ezkalibur,
		"0x498f7bb59c61307de7dea005877220e4406470e9": Ezkalibur,
		"0x240f765af2273b0cab6caff2880d6d8f8b285fa4": Ezkalibur,
		"0x11ef47783740b3f0c9736d54be8ef8953c3ead99": Ezkalibur,
		"0xe67d7d52adb6a14c64b9b1ff8dfec081dd12295e": Ezkalibur,
		"0x6a6d643bdb63bc4d74d166abee57b1be5b3aa9ab": Ezkalibur,

		"0xf8b59f3c3ab33200ec80a8a58b2aa5f5d2a8944c": PancakeSwap,
		"0xa815e2ed7f7d5b0c49fda367f249232a1b9d2883": PancakeSwap,
		"0x5aeaf2883fbf30f3d62471154eda3c0c1b05942d": PancakeSwap,
		"0xd03d8d566183f0086d8d09a84e1e30b58dd5619d": PancakeSwap,

		"0x4476e0c7b72d5ba0dab70b6cc9d0b518c265f01c": BigInt,

		"0xefde2aefe307a7362c7e0e3be019d1491dc7e163": RollupFinance,
		"0xba5f1898f3af43e30b552cac559f11f5d5a2fd1e": RollupFinance,
		"0xf9895dcb26c60743fce3ae54f506c47b763693bb": RollupFinance,
		"0xbbca935fb05d8671b8875293c2d82df73105ac60": RollupFinance,
		"0xc8db02c5900de3ed03139e9ec62a84b2f5e06d7f": RollupFinance,
		"0xc2e9dfb2fded2139b0b60c6781b4c648f1385c16": RollupFinance,
		"0x0ed4a1c66fcdf17341177422e73ca919284e040c": RollupFinance,
		"0x015d01187a77055fadb96c63a9de9f8981e3d0f5": RollupFinance,
		"0x9eb48ce91a77ff37f2edd032c90d7ec6a953b4fd": RollupFinance,
		"0x1f79cce5b8876a07541900e681b31a8e51766b5a": RollupFinance,
		"0xf77ded0ae91c41c432555beb25da2a6658338e5a": RollupFinance,
		"0xcc8b6c37fb97bdc774f0ea7f8b047e8ee057458c": RollupFinance,
		"0xc05c6a229c6ed34b9d31c88421d70ef71fbaabcf": RollupFinance,
		"0xf35eb225803ccf6ffb64be531bb91132a9229196": RollupFinance,
		"0x277e57e33c92c7f145832274384eeef35828db3d": RollupFinance,
		"0x2dddd703578d1f5bd1fad113c9e12f0c416305d5": RollupFinance,
		"0xd23432901c0754e28999661f31ab4aa3633ec2ed": RollupFinance,
		"0x747cc81b6c2b4afe13be59d31591d4229854b826": RollupFinance,
		"0xa0c13b5eca719be98ef27d0c4111cb5abbe97731": RollupFinance,
		"0x824422bd3ba030ca9ca5bcad26dfd4b45409849d": RollupFinance,
		"0xd2589f416c9412572df8d94fa1d58073c9f8f3bf": RollupFinance,
		"0x7bd375b95a1dbcbddb1e8219a01c550b21ed6ea2": RollupFinance,
		"0x5b91962f5eca75e6558e4d32df69b30f75cc6fe5": RollupFinance,
		"0x901b51b9a4214990ac6f0fd50a45df1573b7a51a": RollupFinance,

		"0x1e8f1099a3fe6d2c1a960528394f4feb8f8a288d": Basilisk,
		"0x01541ead71e41d59f315eb2ce3a9441ed7b0a63e": Basilisk,
		"0x1bd39618892115fcf674112b657c3aad1d1b9602": Basilisk,
		"0x827b641c69228a3f259e7904f63f63c405ba6534": Basilisk,
		"0x2f66631fd6e48b21de53c3777d638132675f7c6a": Basilisk,
		"0x4085f99720e699106bc483dab6caed171eda8d15": Basilisk,

		"0x4bba932e9792a2b917d47830c93a9bc79320e4f7": Odos,

		"0xd29aa7bdd3cbb32557973dad995a3219d307721f": Tevaera,
		"0x1eb7bcab5edf75b5e02c9a72d3287e322ebaefdb": Tevaera,
		"0x50b2b7092bcc15fbb8ac74fe9796cf24602897ad": Tevaera,
		"0x0969529a8ea41b47009eb2a590fe71d7942e4f5a": Tevaera,
		"0x9fc20170d613766831f164f1831f4607ae54ff2d": Tevaera,

		"0x4d88434edc8b7ffe215ec598c2290cdc6f58d12d": DraculaFi,

		"0x6dd28C2c5B91DD63b4d4E78EcAC7139878371768": &Service{
			Name: "Merkly",
			Url:  "",
		},

		"0x0000000000000000000000000000000000000000": {
			Name: "WETH",
			Url:  "",
		},
	}

	k := strings.ToLower(a)
	_, ok := m[k]
	if ok {
		return m[k], true
	}

	return nil, false
}
