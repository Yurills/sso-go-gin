
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Link } from "react-router-dom";
import { SSOConfig } from "@/components/auth/SSOConfig";
import { Shield, Users, Zap } from "lucide-react";

const Index = () => {
  return (
    <div className="min-h-screen bg-gradient-to-br from-blue-50 via-white to-purple-50">
      <div className="container mx-auto px-4 py-8">
        <div className="text-center mb-12">
          <h1 className="text-4xl font-bold text-foreground mb-4">
            SSO Authentication Server
          </h1>
          <p className="text-xl text-muted-foreground max-w-2xl mx-auto">
            Secure Single Sign-On portal for seamless authentication across your applications
          </p>
        </div>

        <div className="grid md:grid-cols-3 gap-6 mb-12">
          <Card className="text-center">
            <CardHeader>
              <Shield className="h-12 w-12 text-primary mx-auto mb-4" />
              <CardTitle>Secure Authentication</CardTitle>
            </CardHeader>
            <CardContent>
              <CardDescription>
                OAuth 2.0 compliant authentication with industry-standard security practices
              </CardDescription>
            </CardContent>
          </Card>

          <Card className="text-center">
            <CardHeader>
              <Users className="h-12 w-12 text-primary mx-auto mb-4" />
              <CardTitle>User Management</CardTitle>
            </CardHeader>
            <CardContent>
              <CardDescription>
                Centralized user management with support for multiple client applications
              </CardDescription>
            </CardContent>
          </Card>

          <Card className="text-center">
            <CardHeader>
              <Zap className="h-12 w-12 text-primary mx-auto mb-4" />
              <CardTitle>Easy Integration</CardTitle>
            </CardHeader>
            <CardContent>
              <CardDescription>
                Simple API endpoints for quick integration with your existing applications
              </CardDescription>
            </CardContent>
          </Card>
        </div>

        <div className="text-center mb-12">
          <Button asChild size="lg">
            <Link to="/login">
              Go to Login Portal
            </Link>
          </Button>
        </div>

        {/* <SSOConfig /> */}
      </div>
    </div>
  );
};

export default Index;
