using System;
using Xunit;

namespace SecretHubTest
{
    public class TestSuite
    {
        [Fact]
        public void TestReadSuccess()
        {
            SecretHub.SecretVersion secret = SecretHub.Client.Read("secrethub-xgo/dotnet/test-secret");
            Assert.Equal("super_secret_value", secret.Data);
        }

        [Fact]
        public void TestReadFail()
        {
            string errMessage = "cannot find secret: \"secrethub-xgo/dotnet/not-this-one\": Secret not found (server.secret_not_found) ";
            var ex = Assert.Throws<ApplicationException>(() => SecretHub.Client.Read("secrethub-xgo/dotnet/not-this-one"));
            Assert.Equal(ex.Message, errMessage);
        }

        [Fact]
        public void TestReadSecretSuccess()
        {
            string secret = SecretHub.Client.ReadString("secrethub-xgo/dotnet/test-secret");
            Assert.Equal("super_secret_value", secret);
        }

        [Fact]
        public void TestReadSecretFail()
        {
            string errMessage = "cannot find secret: \"secrethub-xgo/dotnet/not-this-one\": Secret not found (server.secret_not_found) ";
            var ex = Assert.Throws<ApplicationException>(() => SecretHub.Client.ReadString("secrethub-xgo/dotnet/not-this-one"));
            Assert.Equal(ex.Message, errMessage);
        }

        [Fact]
        public void TestResolveSuccess() {
            Assert.Equal("super_secret_value", SecretHub.Client.Resolve("secrethub://secrethub-xgo/dotnet/test-secret"));  
        }

        [Theory]
        [InlineData("secrethub-xgo/dotnet/test-secret", true)]
        [InlineData("secrethub-xgo/dotnet/not-this-one", false)]
        public void TestExists(string path, bool expectedTestResult) {
            Assert.Equal(expectedTestResult, SecretHub.Client.Exists(path));
        }

        [Fact]
        public void TestWriteSuccess() {
            SecretHub.Client.Write("secrethub-xgo/dotnet/new-secret", "new_secret_value");
            String secret = SecretHub.Client.ReadString("secrethub-xgo/dotnet/new-secret");
            Assert.Equal("new_secret_value", secret);
            SecretHub.Client.Remove("secrethub-xgo/dotnet/new-secret");
        }

        [Fact]
        public void TestRemoveSuccess() {
            SecretHub.Client.Write("secrethub-xgo/dotnet/delete-secret", "delete_secret_value");
            Assert.True(SecretHub.Client.Exists("secrethub-xgo/dotnet/delete-secret"));
            SecretHub.Client.Remove("secrethub-xgo/dotnet/delete-secret");
            Assert.False(SecretHub.Client.Exists("secrethub-xgo/dotnet/delete-secret"));
        }

        [Fact]
        public void TestRemoveFail() {
            string errMessage = "Secret not found (server.secret_not_found) ";
            var ex = Assert.Throws<ApplicationException>(() => SecretHub.Client.Remove("secrethub-xgo/dotnet/not-this-one"));
            Assert.Equal(ex.Message, errMessage);
        }
    }
}
