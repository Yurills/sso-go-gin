
import { useState } from "react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Checkbox } from "@/components/ui/checkbox";
import { Eye, EyeOff, Mail, Lock } from "lucide-react";
import { useToast } from "@/hooks/use-toast";
import { ssoApi, type LoginRequest } from "@/services/ssoApi";
import { useSearchParams } from "react-router-dom";

interface LoginFormProps {
  clientApp?: string | null;
  redirectUri?: string | null;
  state?: string | null;
  scope?: string | null;
  responseType?: string | null;
}

export const LoginForm = ({ clientApp, redirectUri, state, scope, responseType }: LoginFormProps) => {
  const [username, setusername] = useState("");
  const [password, setPassword] = useState("");
  const [showPassword, setShowPassword] = useState(false);
  const [rememberMe, setRememberMe] = useState(false);
  const [isLoading, setIsLoading] = useState(false);
  const { toast } = useToast();

  const [searchParams] = useSearchParams();
  const rid = searchParams.get("rid");
  // const csrf_token = searchParams.get("csrf_token")




  function getCSRFTokenFromCookie() {
    const match = RegExp(/csrf_token=([^;]+)/).exec(document.cookie);
    return match ? decodeURIComponent(match[1]) : null;
  }


  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsLoading(true);
    
    try {
      const loginRequest: LoginRequest = {
        username,
        password,
        client_id: clientApp || undefined,
        csrf_ses: getCSRFTokenFromCookie() || undefined,
        rid: rid
      };

      console.log('Attempting login with API:', loginRequest);
      const response = await ssoApi.login(loginRequest);

      if (response.redirect_uri && !response.code) { //handle normal redirect or redirect with code (2FA or continue login)
          window.location.href = response.redirect_uri;
          return;
      }
      
      if (response.code) {
        toast({
          title: "Authentication Successful",
          description: clientApp ? `Redirecting back to ${clientApp}...` : "Welcome to the SSO portal!",
        });

        console.log(response)

        if (response.redirect_uri && response.code && response.state) {
          // SSO flow - redirect back to client application
          const redirectUrl = new URL(response.redirect_uri );
          redirectUrl.searchParams.set('code', response.code);
          redirectUrl.searchParams.set('state', response.state);
          
          console.log('Redirecting to client with auth code:', response.code);
          
          // In production, you would redirect immediately
          window.location.href = redirectUrl.toString();
        } else {
          // Direct login - stay on portal
          console.log("Direct login successful");
        }
      } else {
        toast({
          title: "Authentication Failed",
          description: response.error || "Invalid credentials",
          variant: "destructive",
        });
      }
    } catch (error) {
      console.error('Login error:', error);
      toast({
        title: "Login Error",
        description: "An error occurred during authentication. Please try again.",
        variant: "destructive",
      });
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <form onSubmit={handleSubmit} className="space-y-4">
      <div className="space-y-2">
        <Label htmlFor="username" className="text-sm font-medium">
          Username
        </Label>
        <div className="relative">
          <Mail className="absolute left-3 top-1/2 transform -translate-y-1/2 text-muted-foreground" size={18} />
          <Input
            id="username"
            type="username"
            value={username}
            onChange={(e) => setusername(e.target.value)}
            placeholder="Enter your username"
            className="pl-10 h-12"
            required
            autoCapitalize="off"
          />
        </div>
      </div>

      <div className="space-y-2">
        <Label htmlFor="password" className="text-sm font-medium">
          Password
        </Label>
        <div className="relative">
          <Lock className="absolute left-3 top-1/2 transform -translate-y-1/2 text-muted-foreground" size={18} />
          <Input
            id="password"
            type={showPassword ? "text" : "password"}
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            placeholder="Enter your password"
            className="pl-10 pr-10 h-12"
            required
            autoCapitalize="off"
          />
          <Button
            type="button"
            variant="ghost"
            size="sm"
            className="absolute right-2 top-1/2 transform -translate-y-1/2 h-8 w-8 p-0"
            onClick={() => setShowPassword(!showPassword)}
          >
            {showPassword ? <EyeOff size={16} /> : <Eye size={16} />}
          </Button>
        </div>
      </div>

      <div className="flex items-center justify-between">
        <div className="flex items-center space-x-2">
          <Checkbox
            id="remember"
            checked={rememberMe}
            onCheckedChange={(checked) => setRememberMe(checked as boolean)}
          />
          <Label htmlFor="remember" className="text-sm text-muted-foreground">
            Keep me signed in
          </Label>
        </div>
        <Button variant="link" className="p-0 h-auto text-sm">
          Forgot password?
        </Button>
      </div>

      <Button
        type="submit"
        className="w-full h-12 bg-primary hover:bg-primary/90 transition-colors"
        disabled={isLoading}
      >
        {isLoading ? "Authenticating..." : clientApp ? `Authorize ${clientApp}` : "Sign In"}
      </Button>
    </form>
  );
};
