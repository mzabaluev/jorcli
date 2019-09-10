package jnode

import (
	"bytes"
	"fmt"
	"text/template"
	"time"
)

// block0ConfigTemplate ...
const block0ConfigTemplate = `
{{- with .BlockchainConfiguration -}}
blockchain_configuration:
  block0_date: {{ .Block0Date }}
  discrimination: {{ .Discrimination }}
  block0_consensus: {{ .Block0Consensus }}
  slots_per_epoch: {{ .SlotsPerEpoch }}
  slot_duration: {{ .SlotDuration }}
  kes_update_speed: {{ .KesUpdateSpeed }}
  consensus_genesis_praos_active_slot_coeff: {{ .ConsensusGenesisPraosActiveSlotCoeff }}
  bft_slots_ratio: {{ .BftSlotsRatio }}
  max_number_of_transactions_per_block: {{ .MaxNumberOfTransactionsPerBlock }}
  epoch_stability_depth: {{ .EpochStabilityDepth }}
  {{with .LinearFees -}}
  linear_fees:
    constant: {{ .Constant }}
    coefficient: {{ .Coefficient }}
    certificate: {{ .Certificate }}
  {{- end}}
  consensus_leader_ids:
  {{- range .ConsensusLeaderIds}}
    - {{ . -}}
  {{end}}
{{end}}

{{- with .Initial -}}
initial:
{{- range .}}
  {{- with .Fund}}
  - fund:
      {{- range .}}
      - address: {{ .Address }}
        value: {{ .Value}}
      {{- end -}}
  {{- end -}}
  {{if .Cert}}
  - cert: {{ .Cert }}
  {{- end}}
{{- end}}
{{- end}}
`

// Block0Config Genesis config
type Block0Config struct {
	BlockchainConfiguration BlockchainConfig    // `"blockchain_configuration"`
	Initial                 []BlockchainInitial // `"initial"`
}

// BlockchainConfig ...
type BlockchainConfig struct {
	Discrimination                  string // `"discrimination"`
	Block0Consensus                 string // `"block0_consensus"`
	Block0Date                      int64  // `"block0_date"`
	SlotDuration                    int    // `"slot_duration"`
	SlotsPerEpoch                   int    // `"slots_per_epoch"`
	EpochStabilityDepth             int    // `"epoch_stability_depth"`
	KesUpdateSpeed                  int    // `"kes_update_speed"`
	MaxNumberOfTransactionsPerBlock int    // `"max_number_of_transactions_per_block"`

	BftSlotsRatio                        float64 // `"bft_slots_ratio"`
	ConsensusGenesisPraosActiveSlotCoeff float64 // `"consensus_genesis_praos_active_slot_coeff"`

	LinearFees struct {
		Certificate int // `"certificate"`
		Coefficient int // `"coefficient"`
		Constant    int // `"constant"`
	} // `"linear_fees"`

	ConsensusLeaderIds []string // `"consensus_leader_ids"`
}

// BlockchainInitial ...
type BlockchainInitial struct {
	Fund []InitialFund // `"fund"`
	Cert string        // `"cert"`
}

// InitialFund ...
type InitialFund struct {
	Address string // `"address"`
	Value   uint64 // `"value"`
}

// NewBlock0Config provides default block0 config
func NewBlock0Config() *Block0Config {
	var chainConfig BlockchainConfig

	chainConfig.Discrimination = "test"
	chainConfig.Block0Consensus = "genesis_praos"
	chainConfig.Block0Date = time.Now().Unix()
	chainConfig.SlotDuration = 20
	chainConfig.SlotsPerEpoch = 30
	chainConfig.EpochStabilityDepth = 1
	chainConfig.KesUpdateSpeed = 43200
	chainConfig.BftSlotsRatio = 0.0
	chainConfig.ConsensusGenesisPraosActiveSlotCoeff = 0.1
	chainConfig.MaxNumberOfTransactionsPerBlock = 255
	chainConfig.LinearFees.Certificate = 0
	chainConfig.LinearFees.Coefficient = 0
	chainConfig.LinearFees.Constant = 0

	return &Block0Config{
		BlockchainConfiguration: chainConfig,
	}
}

// ToYaml parses the config template and returns yaml
func (block0Cfg *Block0Config) ToYaml() ([]byte, error) {
	var block0Yaml bytes.Buffer

	t := template.Must(template.New("block0ConfigTemplate").Parse(block0ConfigTemplate))

	err := t.Execute(&block0Yaml, block0Cfg)
	if err != nil {
		return nil, err
	}

	return block0Yaml.Bytes(), nil
}

// AddConsensusLeader to block0 Blockchain Configuration
func (block0Cfg *Block0Config) AddConsensusLeader(leaderPublicKey string) error {
	// FIXME: check validity
	if leaderPublicKey == "" {
		return fmt.Errorf("parameter missing : %s", "leaderPublicKey")
	}

	block0Cfg.BlockchainConfiguration.ConsensusLeaderIds = append(
		block0Cfg.BlockchainConfiguration.ConsensusLeaderIds,
		leaderPublicKey,
	)

	return nil
}

// AddInitialCertificate to block0 Initial config
func (block0Cfg *Block0Config) AddInitialCertificate(cert string) error {
	// FIXME: check validity
	if cert == "" {
		return fmt.Errorf("parameter missing : %s", "cert")
	}

	block0Cfg.Initial = append(
		block0Cfg.Initial,
		BlockchainInitial{Cert: cert},
	)

	return nil
}

// AddInitialFund to block0 Initial config
func (block0Cfg *Block0Config) AddInitialFund(address string, value uint64) error {
	// FIXME: check validity
	if address == "" {
		return fmt.Errorf("parameter missing : %s", "address")
	}

	fundInit := InitialFund{
		Address: address,
		Value:   value,
	}

	block0Cfg.Initial = append(
		block0Cfg.Initial,
		BlockchainInitial{
			Fund: []InitialFund{fundInit},
		},
	)

	return nil
}
