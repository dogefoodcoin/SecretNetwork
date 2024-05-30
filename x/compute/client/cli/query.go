package cli

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"

	// wasmvm "github.com/CosmWasm/wasmvm/v2"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/scrtlabs/SecretNetwork/x/compute/internal/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	wasmUtils "github.com/scrtlabs/SecretNetwork/x/compute/client/utils"
)

func GetQueryCmd() *cobra.Command {
	queryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the wasm module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
		SilenceUsage:               true,
	}
	queryCmd.AddCommand(
		GetCmdListCode(),
		GetCmdListContractByCode(),
		// GetCmdQueryCode(),
		// GetCmdQueryCodeInfo(),
		GetCmdGetContractInfo(),
		// GetCmdGetContractHistory(),
		GetCmdGetContractStateSmart(),
		// GetCmdListPinnedCode(),
		// GetCmdQueryParams(),
		// GetCmdListContractsByCreator(),
	)
	return queryCmd
}

func GetCmdGetContractStateSmart() *cobra.Command {
	decoder := newArgDecoder(asciiDecodeString)
	cmd := &cobra.Command{
		Use:   "query [bech32_address] [query]",
		Short: "Calls contract with given address with query data and prints the returned result",
		Long:  "Calls contract with given address with query data and prints the returned result",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			grpcCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			contractAddr, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return sdkerrors.ErrInvalidAddress.Wrapf("Invalid contract address: %s", args[0])
			}

			queryData, err := decoder.DecodeString(args[1])
			if err != nil {
				return err
			}

			return QueryWithData(contractAddr, queryData, clientCtx, grpcCtx)
		},
		SilenceUsage: true,
	}
	decoder.RegisterFlags(cmd.PersistentFlags(), "key argument")
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetCmdListCode -> gRPC into x/compute/internal/keeper/querier.go: Codes(c context.Context, _ *empty.Empty)
func GetCmdListCode() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list-code",
		Short:   "List all wasm bytecode on the chain",
		Long:    "List all wasm bytecode on the chain",
		Aliases: []string{"list-codes", "codes", "lco"},
		Args:    cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, _ []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.Codes(
				context.Background(),
				&emptypb.Empty{},
			)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
		SilenceUsage: true,
	}
	flags.AddQueryFlagsToCmd(cmd)
	addPaginationFlags(cmd, "list codes")
	return cmd
}

// GetCmdListContractByCode lists all wasm code uploaded for given code id
func GetCmdListContractByCode() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list-contract-by-code [code_id]",
		Short:   "List wasm all bytecode on the chain for given code id",
		Long:    "List wasm all bytecode on the chain for given code id",
		Aliases: []string{"list-contracts-by-code", "list-contracts", "contracts", "lca"},
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			codeID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}
			if codeID == 0 {
				return errors.New("empty code id")
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.ContractsByCodeId(
				context.Background(),
				&types.QueryByCodeIdRequest{
					CodeId: codeID,
				},
			)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
		SilenceUsage: true,
	}
	flags.AddQueryFlagsToCmd(cmd)
	addPaginationFlags(cmd, "list contracts by code")
	return cmd
}

// // GetCmdQueryCode returns the bytecode for a given contract
// func GetCmdQueryCode() *cobra.Command {
// 	cmd := &cobra.Command{
// 		Use:     "code [code_id] [output filename]",
// 		Short:   "Downloads wasm bytecode for given code id",
// 		Long:    "Downloads wasm bytecode for given code id",
// 		Aliases: []string{"source-code", "source"},
// 		Args:    cobra.ExactArgs(2),
// 		RunE: func(cmd *cobra.Command, args []string) error {
// 			clientCtx, err := client.GetClientQueryContext(cmd)
// 			if err != nil {
// 				return err
// 			}

// 			codeID, err := strconv.ParseUint(args[0], 10, 64)
// 			if err != nil {
// 				return err
// 			}

// 			queryClient := types.NewQueryClient(clientCtx)
// 			res, err := queryClient.Code(
// 				context.Background(),
// 				&types.QueryCodeRequest{
// 					CodeId: codeID,
// 				},
// 			)
// 			if err != nil {
// 				return err
// 			}
// 			if len(res.Data) == 0 {
// 				return fmt.Errorf("contract not found")
// 			}

// 			fmt.Printf("Downloading wasm code to %s\n", args[1])
// 			return os.WriteFile(args[1], res.Data, 0o600)
// 		},
// 		SilenceUsage: true,
// 	}
// 	flags.AddQueryFlagsToCmd(cmd)
// 	return cmd
// }

// // GetCmdQueryCodeInfo returns the code info for a given code id
// func GetCmdQueryCodeInfo() *cobra.Command {
// 	cmd := &cobra.Command{
// 		Use:   "code-info [code_id]",
// 		Short: "Prints out metadata of a code id",
// 		Long:  "Prints out metadata of a code id",
// 		Args:  cobra.ExactArgs(1),
// 		RunE: func(cmd *cobra.Command, args []string) error {
// 			clientCtx, err := client.GetClientQueryContext(cmd)
// 			if err != nil {
// 				return err
// 			}

// 			codeID, err := strconv.ParseUint(args[0], 10, 64)
// 			if err != nil {
// 				return err
// 			}

// 			queryClient := types.NewQueryClient(clientCtx)
// 			res, err := queryClient.Code(
// 				context.Background(),
// 				&types.QueryCodeRequest{
// 					CodeId: codeID,
// 				},
// 			)
// 			if err != nil {
// 				return err
// 			}
// 			if res.CodeInfoResponse == nil {
// 				return fmt.Errorf("contract not found")
// 			}

// 			return clientCtx.PrintProto(res.CodeInfoResponse)
// 		},
// 		SilenceUsage: true,
// 	}
// 	flags.AddQueryFlagsToCmd(cmd)
// 	return cmd
// }

// GetCmdGetContractInfo gets details about a given contract
func GetCmdGetContractInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "contract [bech32_address]",
		Short:   "Prints out metadata of a contract given its address",
		Long:    "Prints out metadata of a contract given its address",
		Aliases: []string{"meta", "c"},
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			_, err = sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.ContractInfo(
				context.Background(),
				&types.QueryByContractAddressRequest{
					ContractAddress: args[0],
				},
			)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
		SilenceUsage: true,
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// // GetCmdGetContractState dumps full internal state of a given contract
// func GetCmdGetContractState() *cobra.Command {
// 	cmd := &cobra.Command{
// 		Use:                        "contract-state",
// 		Short:                      "Querying commands for the wasm module",
// 		Aliases:                    []string{"state", "cs", "s"},
// 		DisableFlagParsing:         true,
// 		SuggestionsMinimumDistance: 2,
// 		RunE:                       client.ValidateCmd,
// 		SilenceUsage:               true,
// 	}
// 	cmd.AddCommand(
// 		GetCmdGetContractStateAll(),
// 		GetCmdGetContractStateRaw(),
// 		GetCmdGetContractStateSmart(),
// 	)
// 	return cmd
// }

// func GetCmdGetContractStateAll() *cobra.Command {
// 	cmd := &cobra.Command{
// 		Use:   "all [bech32_address]",
// 		Short: "Prints out all internal state of a contract given its address",
// 		Long:  "Prints out all internal state of a contract given its address",
// 		Args:  cobra.ExactArgs(1),
// 		RunE: func(cmd *cobra.Command, args []string) error {
// 			clientCtx, err := client.GetClientQueryContext(cmd)
// 			if err != nil {
// 				return err
// 			}

// 			_, err = sdk.AccAddressFromBech32(args[0])
// 			if err != nil {
// 				return err
// 			}

// 			pageReq, err := client.ReadPageRequest(withPageKeyDecoded(cmd.Flags()))
// 			if err != nil {
// 				return err
// 			}
// 			queryClient := types.NewQueryClient(clientCtx)
// 			res, err := queryClient.AllContractState(
// 				context.Background(),
// 				&types.QueryAllContractStateRequest{
// 					Address:    args[0],
// 					Pagination: pageReq,
// 				},
// 			)
// 			if err != nil {
// 				return err
// 			}
// 			return clientCtx.PrintProto(res)
// 		},
// 		SilenceUsage: true,
// 	}
// 	flags.AddQueryFlagsToCmd(cmd)
// 	addPaginationFlags(cmd, "contract state")
// 	return cmd
// }

// func GetCmdGetContractStateRaw() *cobra.Command {
// 	decoder := newArgDecoder(hex.DecodeString)
// 	cmd := &cobra.Command{
// 		Use:   "raw [bech32_address] [key]",
// 		Short: "Prints out internal state for key of a contract given its address",
// 		Long:  "Prints out internal state for of a contract given its address",
// 		Args:  cobra.ExactArgs(2),
// 		RunE: func(cmd *cobra.Command, args []string) error {
// 			clientCtx, err := client.GetClientQueryContext(cmd)
// 			if err != nil {
// 				return err
// 			}

// 			_, err = sdk.AccAddressFromBech32(args[0])
// 			if err != nil {
// 				return err
// 			}
// 			queryData, err := decoder.DecodeString(args[1])
// 			if err != nil {
// 				return err
// 			}

// 			queryClient := types.NewQueryClient(clientCtx)
// 			res, err := queryClient.RawContractState(
// 				context.Background(),
// 				&types.QueryRawContractStateRequest{
// 					Address:   args[0],
// 					QueryData: queryData,
// 				},
// 			)
// 			if err != nil {
// 				return err
// 			}
// 			return clientCtx.PrintProto(res)
// 		},
// 		SilenceUsage: true,
// 	}
// 	decoder.RegisterFlags(cmd.PersistentFlags(), "key argument")
// 	flags.AddQueryFlagsToCmd(cmd)
// 	return cmd
// }

// func GetCmdGetContractStateSmart() *cobra.Command {
// 	decoder := newArgDecoder(asciiDecodeString)
// 	cmd := &cobra.Command{
// 		Use:   "smart [bech32_address] [query]",
// 		Short: "Calls contract with given address with query data and prints the returned result",
// 		Long:  "Calls contract with given address with query data and prints the returned result",
// 		Args:  cobra.ExactArgs(2),
// 		RunE: func(cmd *cobra.Command, args []string) error {
// 			clientCtx, err := client.GetClientQueryContext(cmd)
// 			if err != nil {
// 				return err
// 			}

// 			_, err = sdk.AccAddressFromBech32(args[0])
// 			if err != nil {
// 				return err
// 			}
// 			if args[1] == "" {
// 				return errors.New("query data must not be empty")
// 			}

// 			queryData, err := decoder.DecodeString(args[1])
// 			if err != nil {
// 				return fmt.Errorf("decode query: %s", err)
// 			}
// 			if !json.Valid(queryData) {
// 				return errors.New("query data must be json")
// 			}

// 			queryClient := types.NewQueryClient(clientCtx)
// 			res, err := queryClient.SmartContractState(
// 				context.Background(),
// 				&types.QuerySmartContractStateRequest{
// 					Address:   args[0],
// 					QueryData: queryData,
// 				},
// 			)
// 			if err != nil {
// 				return err
// 			}
// 			return clientCtx.PrintProto(res)
// 		},
// 		SilenceUsage: true,
// 	}
// 	decoder.RegisterFlags(cmd.PersistentFlags(), "query argument")
// 	flags.AddQueryFlagsToCmd(cmd)
// 	return cmd
// }

// // GetCmdGetContractHistory prints the code history for a given contract
// func GetCmdGetContractHistory() *cobra.Command {
// 	cmd := &cobra.Command{
// 		Use:     "contract-history [bech32_address]",
// 		Short:   "Prints out the code history for a contract given its address",
// 		Long:    "Prints out the code history for a contract given its address",
// 		Aliases: []string{"history", "hist", "ch"},
// 		Args:    cobra.ExactArgs(1),
// 		RunE: func(cmd *cobra.Command, args []string) error {
// 			clientCtx, err := client.GetClientQueryContext(cmd)
// 			if err != nil {
// 				return err
// 			}

// 			_, err = sdk.AccAddressFromBech32(args[0])
// 			if err != nil {
// 				return err
// 			}

// 			pageReq, err := client.ReadPageRequest(withPageKeyDecoded(cmd.Flags()))
// 			if err != nil {
// 				return err
// 			}
// 			queryClient := types.NewQueryClient(clientCtx)
// 			res, err := queryClient.ContractHistory(
// 				context.Background(),
// 				&types.QueryContractHistoryRequest{
// 					Address:    args[0],
// 					Pagination: pageReq,
// 				},
// 			)
// 			if err != nil {
// 				return err
// 			}

// 			return clientCtx.PrintProto(res)
// 		},
// 		SilenceUsage: true,
// 	}

// 	flags.AddQueryFlagsToCmd(cmd)
// 	addPaginationFlags(cmd, "contract history")
// 	return cmd
// }

// // GetCmdListPinnedCode lists all wasm code ids that are pinned
// func GetCmdListPinnedCode() *cobra.Command {
// 	cmd := &cobra.Command{
// 		Use:   "pinned",
// 		Short: "List all pinned code ids",
// 		Long:  "List all pinned code ids",
// 		Args:  cobra.ExactArgs(0),
// 		RunE: func(cmd *cobra.Command, args []string) error {
// 			clientCtx, err := client.GetClientQueryContext(cmd)
// 			if err != nil {
// 				return err
// 			}

// 			pageReq, err := client.ReadPageRequest(withPageKeyDecoded(cmd.Flags()))
// 			if err != nil {
// 				return err
// 			}
// 			queryClient := types.NewQueryClient(clientCtx)
// 			res, err := queryClient.PinnedCodes(
// 				context.Background(),
// 				&types.QueryPinnedCodesRequest{
// 					Pagination: pageReq,
// 				},
// 			)
// 			if err != nil {
// 				return err
// 			}
// 			return clientCtx.PrintProto(res)
// 		},
// 		SilenceUsage: true,
// 	}
// 	flags.AddQueryFlagsToCmd(cmd)
// 	addPaginationFlags(cmd, "list codes")
// 	return cmd
// }

// // GetCmdListContractsByCreator lists all contracts by creator
// func GetCmdListContractsByCreator() *cobra.Command {
// 	cmd := &cobra.Command{
// 		Use:   "list-contracts-by-creator [creator]",
// 		Short: "List all contracts by creator",
// 		Long:  "List all contracts by creator",
// 		Args:  cobra.ExactArgs(1),
// 		RunE: func(cmd *cobra.Command, args []string) error {
// 			clientCtx, err := client.GetClientQueryContext(cmd)
// 			if err != nil {
// 				return err
// 			}
// 			_, err = sdk.AccAddressFromBech32(args[0])
// 			if err != nil {
// 				return err
// 			}
// 			pageReq, err := client.ReadPageRequest(withPageKeyDecoded(cmd.Flags()))
// 			if err != nil {
// 				return err
// 			}

// 			queryClient := types.NewQueryClient(clientCtx)
// 			res, err := queryClient.ContractsByCreator(
// 				context.Background(),
// 				&types.QueryContractsByCreatorRequest{
// 					CreatorAddress: args[0],
// 					Pagination:     pageReq,
// 				},
// 			)
// 			if err != nil {
// 				return err
// 			}
// 			return clientCtx.PrintProto(res)
// 		},
// 		SilenceUsage: true,
// 	}
// 	flags.AddQueryFlagsToCmd(cmd)
// 	addPaginationFlags(cmd, "list contracts by creator")
// 	return cmd
// }

// type argumentDecoder struct {
// 	// dec is the default decoder
// 	dec                func(string) ([]byte, error)
// 	asciiF, hexF, b64F bool
// }

// func newArgDecoder(def func(string) ([]byte, error)) *argumentDecoder {
// 	return &argumentDecoder{dec: def}
// }

// func (a *argumentDecoder) RegisterFlags(f *flag.FlagSet, argName string) {
// 	f.BoolVar(&a.asciiF, "ascii", false, "ascii encoded "+argName)
// 	f.BoolVar(&a.hexF, "hex", false, "hex encoded "+argName)
// 	f.BoolVar(&a.b64F, "b64", false, "base64 encoded "+argName)
// }

// func (a *argumentDecoder) DecodeString(s string) ([]byte, error) {
// 	found := -1
// 	for i, v := range []*bool{&a.asciiF, &a.hexF, &a.b64F} {
// 		if !*v {
// 			continue
// 		}
// 		if found != -1 {
// 			return nil, errors.New("multiple decoding flags used")
// 		}
// 		found = i
// 	}
// 	switch found {
// 	case 0:
// 		return asciiDecodeString(s)
// 	case 1:
// 		return hex.DecodeString(s)
// 	case 2:
// 		return base64.StdEncoding.DecodeString(s)
// 	default:
// 		return a.dec(s)
// 	}
// }

// func asciiDecodeString(s string) ([]byte, error) {
// 	return []byte(s), nil
// }

// // sdk ReadPageRequest expects binary but we encoded to base64 in our marshaller
// func withPageKeyDecoded(flagSet *flag.FlagSet) *flag.FlagSet {
// 	encoded, err := flagSet.GetString(flags.FlagPageKey)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	raw, err := base64.StdEncoding.DecodeString(encoded)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	err = flagSet.Set(flags.FlagPageKey, string(raw))
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	return flagSet
// }

// // GetCmdQueryParams implements a command to return the current wasm
// // parameters.
// func GetCmdQueryParams() *cobra.Command {
// 	cmd := &cobra.Command{
// 		Use:   "params",
// 		Short: "Query the current wasm parameters",
// 		Args:  cobra.NoArgs,
// 		RunE: func(cmd *cobra.Command, args []string) error {
// 			clientCtx, err := client.GetClientQueryContext(cmd)
// 			if err != nil {
// 				return err
// 			}
// 			queryClient := types.NewQueryClient(clientCtx)

// 			params := &types.QueryParamsRequest{}
// 			res, err := queryClient.Params(cmd.Context(), params)
// 			if err != nil {
// 				return err
// 			}

// 			return clientCtx.PrintProto(&res.Params)
// 		},
// 		SilenceUsage: true,
// 	}

// 	flags.AddQueryFlagsToCmd(cmd)

// 	return cmd
// }

func QueryWithData(contractAddress sdk.AccAddress, queryData []byte, clientCtx client.Context, grpcCtx client.Context) error {
	wasmCtx := wasmUtils.WASMContext{CLIContext: clientCtx}

	codeHash, err := GetCodeHashByContractAddr(clientCtx, contractAddress.String())
	if err != nil {
		return sdkerrors.ErrNotFound.Wrapf("Contract with address %s not found", contractAddress)
	}

	msg := types.SecretMsg{
		CodeHash: codeHash,
		Msg:      queryData,
	}

	queryData, err = wasmCtx.Encrypt(msg.Serialize())
	if err != nil {
		return err
	}
	nonce, _, _, _ := parseEncryptedBlob(queryData) //nolint:dogsled // Ignoring error since we just encrypted it

	queryClient := types.NewQueryClient(grpcCtx)
	res, err := queryClient.QuerySecretContract(
		context.Background(),
		&types.QuerySecretContractRequest{
			ContractAddress: contractAddress.String(),
			Query:           queryData,
		},
	)
	if err != nil {
		return sdkerrors.ErrNotFound.Wrapf("Failed to query secret contract %s. Error: %s", contractAddress, err)
	}

	var resDecrypted []byte
	resDecrypted, err = wasmCtx.Decrypt(res.Data, nonce)
	if err != nil {
		return err
	}
	res.Data = resDecrypted
	decodedResp, err := base64.StdEncoding.DecodeString(string(resDecrypted))
	if err != nil {
		return err
	}

	fmt.Println(string(decodedResp))
	return nil
}

// supports a subset of the SDK pagination params for better resource utilization
func addPaginationFlags(cmd *cobra.Command, query string) {
	cmd.Flags().String(flags.FlagPageKey, "", fmt.Sprintf("pagination page-key of %s to query for", query))
	cmd.Flags().Uint64(flags.FlagLimit, 100, fmt.Sprintf("pagination limit of %s to query for", query))
	cmd.Flags().Bool(flags.FlagReverse, false, "results are sorted in descending order")
}
