package io.secrethub;

import com.sun.jna.*;
import java.util.*;

// SecretHubClient implements a basic client for the SecretHub API.
public class SecretHubClient {  
    static private SecretHub INSTANCE = (SecretHub) Native.load("secrethub", SecretHub.class);
    
    // SecretHub implements the Java interface between Java <-> C <-> Golang
    public interface SecretHub extends Library {
        public String Read(String req);
        public String Write(String req);
    }

    // read returns the secret value at the given path.
    public String read(String path) {
        return INSTANCE.Read("{\"path\":\"" + path + "\"}");
    }

    public void write(String path, byte[] value) throws Exception {
        INSTANCE.Write("{}");
    }
}
