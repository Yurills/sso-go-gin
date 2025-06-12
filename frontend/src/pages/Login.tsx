
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { LoginForm } from "@/components/auth/LoginForm";
import { useEffect, useState } from "react";
import { Button } from "@/components/ui/button";
import { ArrowLeft, AlertCircle } from "lucide-react";
import { ssoApi, type AuthorizeRequest } from "@/services/ssoApi";
import { Alert, AlertDescription } from "@/components/ui/alert";
import { useSearchParams } from "react-router-dom";

const Login = () => {
  const [clientApp, setClientApp] = useState<string | null>(null);
  const [redirectUri, setRedirectUri] = useState<string | null>(null);
  const [state, setState] = useState<string | null>(null);
  const [scope, setScope] = useState<string | null>(null);
  const [responseType, setResponseType] = useState<string | null>(null);
  const [authError, setAuthError] = useState<string | null>(null);
  const [isValidating, setIsValidating] = useState(false);
  const [isValidRequest, setIsValidRequest] = useState(false);
  const [searchParams] = useSearchParams();

  useEffect(() => {
      const clientId = searchParams.get('client_id');
      setClientApp(clientId)



  },[])
  //   const validateAuthRequest = async () => {
  //     const urlParams = new URLSearchParams(window.location.search);
  //     const clientId = "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11";
  //     const redirect = urlParams.get('redirect_uri');
  //     const stateParam = urlParams.get('state');
  //     const scopeParam = urlParams.get('scope');
  //     const responseTypeParam = urlParams.get('response_type') || 'code';
      
  //     setClientApp(clientId);
  //     setRedirectUri(redirect);
  //     setState(stateParam);
  //     setScope(scopeParam);
  //     setResponseType(responseTypeParam);
      
  //     console.log('SSO Request received:', { clientId, redirect, stateParam, scopeParam, responseTypeParam });

  //     // If we have the required parameters, validate with the API
  //     // if (clientId && redirect) {
  //     //   setIsValidating(true);
        
  //     //   const authRequest: AuthorizeRequest = {
  //     //     client_id: clientId,
  //     //     redirect_uri: redirect,
  //     //     response_type: responseTypeParam,
  //     //     scope: scopeParam || undefined,
  //     //     state: stateParam || undefined,
  //     //   };

  //     //   try {
  //     //     const validation = await ssoApi.authorize(authRequest);
          
  //     //     if (validation.valid) {
  //     //       setIsValidRequest(true);
  //     //       setAuthError(null);
  //     //     } else {
  //     //       setAuthError(validation.error || 'Invalid authorization request');
  //     //       setIsValidRequest(false);
  //     //     }
  //     //   } catch (error) {
  //     //     console.error('Authorization validation failed:', error);
  //     //     setAuthError('Failed to validate authorization request');
  //     //     setIsValidRequest(false);
  //     //   } finally {
  //     //     setIsValidating(false);
  //     //   }
  //     // } else if (clientId || redirect) {
  //     //   // Some parameters present but not all required ones
  //     //   setAuthError('Missing required parameters: client_id and redirect_uri');
  //     // } else {
  //     //   // No SSO parameters - direct access
  //     //   setIsValidRequest(true);
  //     // }
  //   };

  //   // validateAuthRequest();
  // }, []);

  const handleCancel = () => {
    if (redirectUri) {
      const errorUrl = new URL(redirectUri);
      errorUrl.searchParams.set('error', 'access_denied');
      errorUrl.searchParams.set('error_description', 'User cancelled authorization');
      if (state) errorUrl.searchParams.set('state', state);
      
      window.location.href = errorUrl.toString();
    } else {
      window.history.back();
    }
  };

  if (isValidating) {
    return (
      <div className="min-h-screen bg-gradient-to-br from-blue-50 via-white to-purple-50 flex items-center justify-center p-4">
        <Card className="shadow-2xl border-0 bg-white/80 backdrop-blur-sm">
          <CardContent className="p-8 text-center">
            <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-primary mx-auto mb-4"></div>
            <p>Validating authorization request...</p>
          </CardContent>
        </Card>
      </div>
    );
  }

  if (authError) {
    return (
      <div className="min-h-screen bg-gradient-to-br from-blue-50 via-white to-purple-50 flex items-center justify-center p-4">
        <Card className="shadow-2xl border-0 bg-white/80 backdrop-blur-sm max-w-md">
          <CardHeader className="text-center">
            <CardTitle className="text-2xl font-bold text-destructive">Authorization Error</CardTitle>
          </CardHeader>
          <CardContent className="space-y-4">
            <Alert variant="destructive">
              <AlertCircle className="h-4 w-4" />
              <AlertDescription>{authError}</AlertDescription>
            </Alert>
            <div className="flex justify-center">
              <Button onClick={handleCancel} variant="outline">
                <ArrowLeft size={16} className="mr-2" />
                Go Back
              </Button>
            </div>
          </CardContent>
        </Card>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-gradient-to-br from-blue-50 via-white to-purple-50 flex items-center justify-center p-4">
      <div className="w-full max-w-md">
        <Card className="shadow-2xl border-0 bg-white/80 backdrop-blur-sm">
          <CardHeader className="text-center space-y-2 pb-4">
            <CardTitle className="text-2xl font-bold text-foreground">
              SSO Authentication
            </CardTitle>
            <CardDescription className="text-muted-foreground">
              {clientApp ? (
                <>
                  <strong>{clientApp}</strong> is requesting access to your account
                </>
              ) : (
                "Sign in to continue"
              )}
            </CardDescription>
          </CardHeader>
          <CardContent className="space-y-6">
            {clientApp && (
              <div className="bg-muted/50 p-4 rounded-lg border text-sm">
                <div className="flex justify-between items-center mb-2">
                  <span className="font-medium">Client Application:</span>
                  <span className="text-muted-foreground">{clientApp}</span>
                </div>
                {redirectUri && (
                  <div className="flex justify-between items-center mb-2">
                    <span className="font-medium">Redirect URI:</span>
                    <span className="text-muted-foreground text-xs truncate max-w-[200px]">
                      {redirectUri}
                    </span>
                  </div>
                )}
                {scope && (
                  <div className="flex justify-between items-center">
                    <span className="font-medium">Requested Scope:</span>
                    <span className="text-muted-foreground text-xs">{scope}</span>
                  </div>
                )}
              </div>
            )}
            
            <LoginForm 
              clientApp={clientApp} 
              redirectUri={redirectUri} 
              state={state}
              scope={scope}
              responseType={responseType}
            />
            
            {clientApp && (
              <div className="flex justify-center">
                <Button
                  variant="ghost"
                  onClick={handleCancel}
                  className="text-sm text-muted-foreground hover:text-foreground"
                >
                  <ArrowLeft size={16} className="mr-2" />
                  Cancel and go back
                </Button>
              </div>
            )}
          </CardContent>
        </Card>
        
        <div className="mt-6 text-center text-sm text-muted-foreground">
          <p>
            Secure SSO Authentication Portal
          </p>
        </div>
      </div>
    </div>
  );
};

export default Login;
