using System;
using Xunit;
using System.Text.RegularExpressions;

namespace SecretHubTest
{
    public class TestSuite
    {
        [Fact]
        public void TestReadSuccess()
        {
            SecretHub.SecretVersion expectedSecret = new SecretHub.SecretVersion();
            SecretHub.SecretVersion secret = SecretHub.Client.Read("secrethub-xgo/dotnet/test-secret");
            Assert.Equal(new Guid("a2628f70-dade-49b4-b4db-eca16c15e1d2"), secret.SecretVersionID);
            Assert.Equal(new Guid("5b6a82f7-1b55-4e23-ac76-0a4f1d2fa826"), secret.Secret.SecretID);
            Assert.Equal(2, secret.Version);
            Assert.Equal("super_secret_value", secret.Data);
            Assert.Equal(DateTime.Parse("9/2/2020 2:11:49 PM",
                          System.Globalization.CultureInfo.InvariantCulture), secret.CreatedAt);
            Assert.Equal(DateTime.Parse("8/31/2020 2:39:50 PM",
                          System.Globalization.CultureInfo.InvariantCulture), secret.Secret.CreatedAt);
            Assert.Equal("ok", secret.Status);
        }

        [Fact]
        public void TestReadFail()
        {
            Regex expectedErrorRegex = new Regex(@"^.*\(server\.secret_not_found\) $");
            var ex = Assert.Throws<ApplicationException>(() => SecretHub.Client.Read("secrethub-xgo/dotnet/not-this-one"));
            Assert.True(expectedErrorRegex.IsMatch(ex.Message), "error should end in the (server.secret_not_found) error code");
        }

        [Fact]
        public void TestReadStringSuccess()
        {
            string secret = SecretHub.Client.ReadString("secrethub-xgo/dotnet/test-secret");
            Assert.Equal("super_secret_value", secret);
        }

        [Fact]
        public void TestReadStringFail()
        {
            Regex expectedErrorRegex = new Regex(@"^.*\(server\.secret_not_found\) $");
            var ex = Assert.Throws<ApplicationException>(() => SecretHub.Client.ReadString("secrethub-xgo/dotnet/not-this-one"));
            Assert.True(expectedErrorRegex.IsMatch(ex.Message), "error should end in the (server.secret_not_found) error code");
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
        public void TestExistsException()
        {
            Regex expectedErrorRegex = new Regex(@"^.*\(api\.invalid_secret_path\) $");
            var ex = Assert.Throws<ApplicationException>(() => SecretHub.Client.Exists("not-a-path"));
            Assert.True(expectedErrorRegex.IsMatch(ex.Message), "error should end in the (api.invalid_secret_path) error code");
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
            Regex expectedErrorRegex = new Regex(@"^.*\(server\.secret_not_found\) $");
            var ex = Assert.Throws<ApplicationException>(() => SecretHub.Client.Remove("secrethub-xgo/dotnet/not-this-one"));
            Assert.True(expectedErrorRegex.IsMatch(ex.Message), "error should end in the (server.secret_not_found) error code");
        }
    }
}
