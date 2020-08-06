# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from registry import registry_pb2 as registry_dot_registry__pb2


class RegistryStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.GetService = channel.unary_unary(
                '/go.micro.registry.Registry/GetService',
                request_serializer=registry_dot_registry__pb2.GetRequest.SerializeToString,
                response_deserializer=registry_dot_registry__pb2.GetResponse.FromString,
                )
        self.Register = channel.unary_unary(
                '/go.micro.registry.Registry/Register',
                request_serializer=registry_dot_registry__pb2.Service.SerializeToString,
                response_deserializer=registry_dot_registry__pb2.EmptyResponse.FromString,
                )
        self.Deregister = channel.unary_unary(
                '/go.micro.registry.Registry/Deregister',
                request_serializer=registry_dot_registry__pb2.Service.SerializeToString,
                response_deserializer=registry_dot_registry__pb2.EmptyResponse.FromString,
                )
        self.ListServices = channel.unary_unary(
                '/go.micro.registry.Registry/ListServices',
                request_serializer=registry_dot_registry__pb2.ListRequest.SerializeToString,
                response_deserializer=registry_dot_registry__pb2.ListResponse.FromString,
                )
        self.Watch = channel.unary_stream(
                '/go.micro.registry.Registry/Watch',
                request_serializer=registry_dot_registry__pb2.WatchRequest.SerializeToString,
                response_deserializer=registry_dot_registry__pb2.Result.FromString,
                )


class RegistryServicer(object):
    """Missing associated documentation comment in .proto file."""

    def GetService(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Register(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Deregister(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ListServices(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Watch(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_RegistryServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'GetService': grpc.unary_unary_rpc_method_handler(
                    servicer.GetService,
                    request_deserializer=registry_dot_registry__pb2.GetRequest.FromString,
                    response_serializer=registry_dot_registry__pb2.GetResponse.SerializeToString,
            ),
            'Register': grpc.unary_unary_rpc_method_handler(
                    servicer.Register,
                    request_deserializer=registry_dot_registry__pb2.Service.FromString,
                    response_serializer=registry_dot_registry__pb2.EmptyResponse.SerializeToString,
            ),
            'Deregister': grpc.unary_unary_rpc_method_handler(
                    servicer.Deregister,
                    request_deserializer=registry_dot_registry__pb2.Service.FromString,
                    response_serializer=registry_dot_registry__pb2.EmptyResponse.SerializeToString,
            ),
            'ListServices': grpc.unary_unary_rpc_method_handler(
                    servicer.ListServices,
                    request_deserializer=registry_dot_registry__pb2.ListRequest.FromString,
                    response_serializer=registry_dot_registry__pb2.ListResponse.SerializeToString,
            ),
            'Watch': grpc.unary_stream_rpc_method_handler(
                    servicer.Watch,
                    request_deserializer=registry_dot_registry__pb2.WatchRequest.FromString,
                    response_serializer=registry_dot_registry__pb2.Result.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'go.micro.registry.Registry', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class Registry(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def GetService(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/go.micro.registry.Registry/GetService',
            registry_dot_registry__pb2.GetRequest.SerializeToString,
            registry_dot_registry__pb2.GetResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Register(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/go.micro.registry.Registry/Register',
            registry_dot_registry__pb2.Service.SerializeToString,
            registry_dot_registry__pb2.EmptyResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Deregister(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/go.micro.registry.Registry/Deregister',
            registry_dot_registry__pb2.Service.SerializeToString,
            registry_dot_registry__pb2.EmptyResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ListServices(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/go.micro.registry.Registry/ListServices',
            registry_dot_registry__pb2.ListRequest.SerializeToString,
            registry_dot_registry__pb2.ListResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Watch(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_stream(request, target, '/go.micro.registry.Registry/Watch',
            registry_dot_registry__pb2.WatchRequest.SerializeToString,
            registry_dot_registry__pb2.Result.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
