package io.secrethub;

import com.sun.jna.*;
import java.util.*;
import io.secrethub.proto.SecretHub.ReadRequest;
import io.secrethub.proto.SecretHub.ReadResponse;
import io.secrethub.proto.SecretHub.Secret;

// SecretHubClient implements a basic client for the SecretHub API.
public class SecretHubClient {  
    static private SecretHub INSTANCE = (SecretHub) Native.load("secrethub", SecretHub.class);
    
    // SecretHub implements the Java interface between Java <-> C <-> Golang
    public interface SecretHub extends Library {
        public String Read(String req);
        public String Write(String req);
    }

    // read returns the secret value at the given path.
    public Secret read(String path) throws Exception {
        var req = ReadRequest.newBuilder().setPath(path).build();
        var encoded = Base64.getEncoder().encodeToString(req.toByteArray());

        var payload = INSTANCE.Read(encoded);
        var dec = Base64.getDecoder().decode(payload);
        var res = ReadResponse.parseFrom(dec);

        return res.getSecret();
    }

    public void write(String path, byte[] value) throws Exception {
        INSTANCE.Write("{}");
    }
}
