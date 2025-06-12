
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Code, Server, Key } from "lucide-react";

export const SSOConfig = () => {
  const apiEndpoints = `# SSO Server API Endpoints

## 1. Authorization Endpoint: POST /authorize
Validates the authorization request parameters.

Request Body:
{
  "client_id": "your-client-app-id",
  "redirect_uri": "https://client-app.com/auth/callback",
  "response_type": "code",
  "scope": "openid profile email", // optional
  "state": "random-state-value"    // optional
}

Response:
{
  "valid": true
}
// or
{
  "valid": false,
  "error": "Invalid client_id"
}

## 2. Login Endpoint: POST /login
Authenticates user and returns authorization code.

Request Body:
{
  "email": "user@example.com",
  "password": "userpassword",
  "client_id": "your-client-app-id",     // optional for SSO flow
  "redirect_uri": "https://client-app.com/auth/callback", // optional
  "state": "random-state-value"          // optional
}

Response:
{
  "success": true,
  "authorization_code": "auth_code_12345",
  "user": {
    "id": "user123",
    "email": "user@example.com",
    "name": "John Doe"
  }
}
// or
{
  "success": false,
  "error": "Invalid credentials"
}

## 3. Token Endpoint: POST /token
Exchanges authorization code for access token.

Request Body:
{
  "grant_type": "authorization_code",
  "code": "auth_code_12345",
  "redirect_uri": "https://client-app.com/auth/callback",
  "client_id": "your-client-app-id",
  "client_secret": "your-client-secret"
}

Response:
{
  "access_token": "access_token_67890",
  "token_type": "Bearer",
  "expires_in": 3600,
  "refresh_token": "refresh_token_abcde", // optional
  "scope": "openid profile email"
}

## 4. User Info Endpoint: GET /userinfo
Returns user information using access token.

Headers:
Authorization: Bearer access_token_67890

Response:
{
  "sub": "user123",
  "email": "user@example.com",
  "name": "John Doe",
  "picture": "https://example.com/avatar.jpg"
}`;

  const envConfig = `# Environment Configuration

# Add to your .env file:
VITE_SSO_API_URL=http://localhost:3000

# For production:
VITE_SSO_API_URL=https://your-sso-api.com`;

  return (
    <div className="space-y-6">
      <Card>
        <CardHeader>
          <CardTitle className="flex items-center gap-2">
            <Server size={20} />
            API Endpoints Implementation
          </CardTitle>
          <CardDescription>
            Implement these endpoints in your SSO server backend
          </CardDescription>
        </CardHeader>
        <CardContent>
          <pre className="bg-muted p-4 rounded-lg text-sm overflow-x-auto whitespace-pre-wrap">
            <code>{apiEndpoints}</code>
          </pre>
        </CardContent>
      </Card>

      <Card>
        <CardHeader>
          <CardTitle className="flex items-center gap-2">
            <Key size={20} />
            Environment Configuration
          </CardTitle>
          <CardDescription>
            Configure your API base URL
          </CardDescription>
        </CardHeader>
        <CardContent>
          <pre className="bg-muted p-4 rounded-lg text-sm overflow-x-auto">
            <code>{envConfig}</code>
          </pre>
        </CardContent>
      </Card>

      <Card>
        <CardHeader>
          <CardTitle className="flex items-center gap-2">
            <Code size={20} />
            OAuth 2.0 Flow Summary
          </CardTitle>
        </CardHeader>
        <CardContent className="space-y-3 text-sm">
          <div className="flex items-start gap-3">
            <span className="bg-primary text-primary-foreground rounded-full w-6 h-6 flex items-center justify-center text-xs font-bold">1</span>
            <div>
              <strong>Authorization Request:</strong> Client redirects user to your portal with client_id, redirect_uri, etc.
            </div>
          </div>
          <div className="flex items-start gap-3">
            <span className="bg-primary text-primary-foreground rounded-full w-6 h-6 flex items-center justify-center text-xs font-bold">2</span>
            <div>
              <strong>Validation:</strong> Portal validates the request via /authorize endpoint
            </div>
          </div>
          <div className="flex items-start gap-3">
            <span className="bg-primary text-primary-foreground rounded-full w-6 h-6 flex items-center justify-center text-xs font-bold">3</span>
            <div>
              <strong>User Authentication:</strong> User logs in via /login endpoint, receives authorization code
            </div>
          </div>
          <div className="flex items-start gap-3">
            <span className="bg-primary text-primary-foreground rounded-full w-6 h-6 flex items-center justify-center text-xs font-bold">4</span>
            <div>
              <strong>Code Exchange:</strong> Client exchanges authorization code for access token via /token endpoint
            </div>
          </div>
          <div className="flex items-start gap-3">
            <span className="bg-primary text-primary-foreground rounded-full w-6 h-6 flex items-center justify-center text-xs font-bold">5</span>
            <div>
              <strong>User Info:</strong> Client gets user information via /userinfo endpoint using access token
            </div>
          </div>
        </CardContent>
      </Card>
    </div>
  );
};
