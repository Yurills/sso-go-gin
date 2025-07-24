import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { ssoApi } from "@/services/ssoApi";
import { useState } from "react";

const TwoFA = () => {
    const [loading, setLoading] = useState(false);
    
    const handleVerify = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        setLoading(true);
        try {
            const response = await ssoApi.continue_login();
            if (response.code && response.redirect_uri && response.state) {
                window.location.href = `${response.redirect_uri}?code=${response.code}&state=${response.state}`;
            }
        } catch (error) {
            console.error("Error during 2FA verification:", error);
        } finally {
            setLoading(false);
        }
    };

    return (
        <div className="min-h-screen flex items-center justify-center bg-gradient-to-br from-blue-50 via-white to-purple-50">
            <form 
                onSubmit={handleVerify} 
                className="w-full max-w-md bg-white shadow-xl rounded-2xl p-8 space-y-6 border border-gray-100"
            >
                <div className="text-center">
                    <h1 className="text-3xl font-semibold text-gray-900">Two-Factor Authentication</h1>
                    <p className="mt-2 text-sm text-gray-600">Enter the code sent to your device to continue.</p>
                </div>

                <Input 
                    className="text-center text-lg tracking-widest"
                    type="text"
                    placeholder="123 456"
                    autoCapitalize="off"
                    required
                />

                <Button 
                    type="submit" 
                    className="w-full h-12 text-base font-medium"
                    disabled={loading}
                >
                    {loading ? "Verifying..." : "Verify"}
                </Button>
            </form>
        </div>
    );
};

export default TwoFA;
