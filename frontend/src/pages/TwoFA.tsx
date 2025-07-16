import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { ssoApi } from "@/services/ssoApi";
import { useSearchParams } from "react-router-dom";

const TwoFA = () => {

    const handleVerify = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        console.log('Verifying 2FA code...');
        try {
            const response = await ssoApi.continue_login();
            if (response.code && response.redirect_uri && response.state) {
                // Redirect to the specified URI with the code
                window.location.href = `${response.redirect_uri}?code=${response.code}&state=${response.state}`;
            }
        } catch (error) {
            console.error('Error during 2FA verification:', error); 
        }
    }

    return (
        <form onSubmit={handleVerify}>
        <div className="min-h-screen bg-gradient-to-br from-blue-50 via-white to-purple-50">
        <div className="container mx-auto px-4 py-8">
            <div className="text-center mb-12">
                <h1 className="text-4xl font-bold text-foreground mx-4 my-2">Two-Factor Authentication</h1>
                <p className="text-xl text-muted-foreground max-w-2xl mx-auto">Please enter the code sent to your device.</p>
                <Input className="text-center" type="text" placeholder="Enter code" autoCapitalize="off"/>
                <div className="text-center mb-12">
                    <Button type="submit" className="h-12 bg-primary hover:bg-primary/90 transition-colors">
                        <span>Verify</span>
                    </Button>

                </div>
            </div>
        </div>
        </div>
        </form>
    );
};

export default TwoFA;