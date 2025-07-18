import { useState } from "react";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Button } from "@/components/ui/button";
import { Textarea } from "@/components/ui/textarea";
import { Switch } from "@/components/ui/switch";
import { ClipboardCopyIcon } from "lucide-react";

const RegisterAuthClient = () => {
  const [form, setForm] = useState({
    name: "",
    description: "",
    auth_redirect_callback_uri: "",
    sso_redirect_callback_uri: "",
    scope: "",
    active: true,
    config_profile: "{}",
    private_key: "",
    public_key: "",
  });

  const [generated, setGenerated] = useState<null | {
    client_id: string;
    client_secret: string;
  }>(null);

  const [loading, setLoading] = useState(false);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    const { name, value } = e.target;
    setForm({ ...form, [name]: value });
  };

  const handleSwitch = (checked: boolean) => {
    setForm({ ...form, active: checked });
  };

  const handleCopy = (text: string) => {
    navigator.clipboard.writeText(text);
    alert("Copied to clipboard");
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);

    try {
      const res = await fetch("/api/admin/auth-clients", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(form),
      });

      if (!res.ok) throw new Error("Registration failed");

      const data = await res.json(); // Expecting: { client_id, client_secret }
      setGenerated({
        client_id: data.client_id,
        client_secret: data.client_secret,
      });

      setForm({
        name: "",
        description: "",
        auth_redirect_callback_uri: "",
        sso_redirect_callback_uri: "",
        scope: "",
        active: true,
        config_profile: "{}",
        private_key: "",
        public_key: "",
      });
    } catch (error) {
      console.error(error);
      alert("Error registering client.");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="min-h-screen bg-gray-50 flex items-center justify-center px-4">
      <form 
        onSubmit={handleSubmit}
        className="bg-white shadow-md border border-gray-200 rounded-xl w-full max-w-2xl p-8 space-y-6"
      >
        <h1 className="text-2xl font-semibold text-center text-gray-900">
          Register New SSO Auth Client
        </h1>

        {/* Admin-typed fields */}
        {["name", "description", "auth_redirect_callback_uri", "sso_redirect_callback_uri", "scope"].map(field => (
          <div key={field}>
            <Label htmlFor={field}>{field.replace(/_/g, " ").replace(/\b\w/g, l => l.toUpperCase())}</Label>
            <Input name={field} value={(form as any)[field]} onChange={handleChange} required={field !== "scope"} />
          </div>
        ))}

        {/* Active Toggle */}
        <div className="flex items-center justify-between">
          <Label htmlFor="active">Active</Label>
          <Switch checked={form.active} onCheckedChange={handleSwitch} />
        </div>

        {/* Config Profile */}
        <div>
          <Label htmlFor="config_profile">Config Profile (JSON)</Label>
          <Textarea name="config_profile" value={form.config_profile} onChange={handleChange} />
        </div>

        {/* Keys */}
        <div>
          <Label htmlFor="private_key">Private Key</Label>
          <Textarea name="private_key" value={form.private_key} onChange={handleChange} required />
        </div>

        <div>
          <Label htmlFor="public_key">Public Key</Label>
          <Textarea name="public_key" value={form.public_key} onChange={handleChange} required />
        </div>

        <Button type="submit" className="w-full h-12" disabled={loading}>
          {loading ? "Registering..." : "Register Client"}
        </Button>

        {/* Display generated credentials */}
        {generated && (
          <div className="mt-6 bg-gray-100 rounded-lg p-4 border text-sm">
            <p className="mb-2 font-medium text-gray-700">🎉 Client registered successfully!</p>
            <p className="text-red-600 font-semibold">⚠️ Please save these credentials. You won’t see them again.</p>
            <div className="mt-4 space-y-2">
              <div className="flex items-center justify-between">
                <span className="font-mono text-xs break-all">Client ID: {generated.client_id}</span>
                <Button variant="ghost" size="sm" onClick={() => handleCopy(generated.client_id)}>
                  <ClipboardCopyIcon className="h-4 w-4" />
                </Button>
              </div>
              <div className="flex items-center justify-between">
                <span className="font-mono text-xs break-all">Client Secret: {generated.client_secret}</span>
                <Button variant="ghost" size="sm" onClick={() => handleCopy(generated.client_secret)}>
                  <ClipboardCopyIcon className="h-4 w-4" />
                </Button>
              </div>
            </div>
          </div>
        )}
      </form>
    </div>
  );
};

export default RegisterAuthClient;
