package cc

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "CC/api/cc/cc"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod:      "ShowExchange",
					Use:            "show-exchange [id]",
					Short:          "Query show-exchange",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},

				{
					RpcMethod:      "ListExchange",
					Use:            "list-exchange",
					Short:          "Query list-exchange",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
				},

				{
					RpcMethod:      "ListSenderExchange",
					Use:            "list-sender-exchange [sender]",
					Short:          "Query listSender-exchange",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "sender"}},
				},

				{
					RpcMethod:      "ListSenderExchange",
					Use:            "list-sender-exchange [sender]",
					Short:          "Query listSender-exchange",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "sender"}},
				},

				{
					RpcMethod:      "SenderlistExchange",
					Use:            "senderlist-exchange [sender]",
					Short:          "Query senderlist-exchange",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "sender"}},
				},

				{
					RpcMethod:      "DatelistExchange",
					Use:            "datelist-exchange [sender]",
					Short:          "Query datelist-exchange",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "sender"}},
				},

				{
					RpcMethod:      "DatelistExchange",
					Use:            "datelist-exchange [date]",
					Short:          "Query Datelist-exchange",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "date"}},
				},

				{
					RpcMethod:      "PleasetExchange",
					Use:            "pleaset-exchange [date]",
					Short:          "Query pleaset-exchange",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "date"}},
				},

				{
					RpcMethod:      "DatelistExchange",
					Use:            "datelist-exchange [date]",
					Short:          "Query datelist-exchange",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "date"}},
				},

				{
					RpcMethod:      "SenderlistExchange",
					Use:            "senderlist-exchange [sender]",
					Short:          "Query senderlist-exchange",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "sender"}},
				},

				{
					RpcMethod:      "RecieverlistExchange",
					Use:            "recieverlist-exchange [reciever]",
					Short:          "Query recieverlist-exchange",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "reciever"}},
				},

				{
					RpcMethod:      "ListSenderExchange",
					Use:            "list-sender-exchange [sender]",
					Short:          "Query listSenderExchange",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "sender"}},
				},

				{
					RpcMethod:      "ListSenderExchange",
					Use:            "list-sender-exchange [sender]",
					Short:          "Query listSenderExchange",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "sender"}},
				},

				{
					RpcMethod:      "ListDateExchange",
					Use:            "list-date-exchange [date]",
					Short:          "Query listDateExchange",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "date"}},
				},

				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreateExchange",
					Use:            "create-exchange [date] [amount] [message] [reciever]",
					Short:          "Send a create-exchange tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "date"}, {ProtoField: "amount"}, {ProtoField: "message"}, {ProtoField: "reciever"}},
				},
				{
					RpcMethod:      "UpdateExchange",
					Use:            "update-exchange [date] [message] [id]",
					Short:          "Send a update-exchange tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "date"}, {ProtoField: "message"}, {ProtoField: "id"}},
				},
				{
					RpcMethod:      "DeleteExchange",
					Use:            "delete-exchange [id]",
					Short:          "Send a delete-exchange tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
