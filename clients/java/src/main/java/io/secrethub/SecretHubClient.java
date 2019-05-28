package io.secrethub;

import com.sun.jna.*;
import java.util.*;

// SecretHubClient implements a basic client for the SecretHub API.
public class SecretHubClient {  
    static private SecretHub INSTANCE = (SecretHub) Native.load("secrethub", SecretHub.class);
    
    // SecretHub implements the Java interface between Java <-> C <-> Golang
    public interface SecretHub extends Library {
        public String Read(String credential, String passphrase, String path);
        public boolean Exists(String credential, String passphrase, String path);
    }

    private String credential;
    private String passphrase;

    // SecretHubClient initializes the client, sourcing configuration
    // from the environment. 
    public SecretHubClient() {
        this.credential = System.getenv("SECRETHUB_CREDENTIAL");
        this.passphrase = System.getenv("SECRETHUB_CREDENTIAL_PASSPHRASE");
    }

    // SecretHubClient initializes the client with the given values.
    public SecretHubClient(String credential, String passphrase) {
        this.credential = credential;
        this.passphrase = passphrase;
    }

    // read returns the secret value at the given path.
    public String read(String path) {
        return INSTANCE.Read(this.credential, this.passphrase, path);
    }

    // exists returns true when a secret exists at the given path.
    public boolean exists(String path) {
        return INSTANCE.Exists(this.credential, this.passphrase, path);
    }
}