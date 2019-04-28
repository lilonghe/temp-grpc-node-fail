function main() {
    var PROTO_PATH = __dirname + '/src/proto/user.proto';
    var grpc = require('grpc');
    var protoLoader = require('@grpc/proto-loader');
    // Suggested options for similarity to existing grpc.load behavior
    var packageDefinition = protoLoader.loadSync(
        PROTO_PATH,
        {keepCase: true,
        longs: String,
        enums: String,
        defaults: true,
        oneofs: true
        });
        
    var protoDescriptor = grpc.loadPackageDefinition(packageDefinition);
    // The protoDescriptor object has the full package hierarchy
    var UserService = protoDescriptor.UserService;
    var server = new grpc.Server();

    server.addService(UserService.service, {
        GetUserById
    })

    server.bind(":8091", grpc.ServerCredentials.createInsecure());
    console.log('listen -> 8091');
    server.start();
}

function GetUserById(call, callback) {
    callback(null, {id: 1})
}

main();