import { useState } from "react";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Button } from "@/components/ui/button";
import { Textarea } from "@/components/ui/textarea";
import { Switch } from "@/components/ui/switch";
import { set } from "date-fns";

const RegisterAuthClient = () => {
  const [form, setForm] = useState({
    name: "",
    client_id: "",
    client_secret: "",
    description: "",
    auth_redirect_callback_uri: "",
    sso_redirect_callback_uri: "",
    scope: "",
    active: true,
    config_profile: "{}",
    private_key: "",
    public_key: "",
  });

  const [loading, setLoading] = useState(false);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    const { name, value } = e.target;
    setForm({ ...form, [name]: value });
  };

  const handleSwitch = (checked: boolean) => {
    setForm({ ...form, active: checked });
  };
  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);

    try {
      const res = await fetch("/api/admin/register-client", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(form),
      });

      if (!res.ok) throw new Error("Registration failed");

      alert("Client registered successfully!");

      setForm({
        name: "",
        client_id: "",
        client_secret: "",
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
    <div className="min-h-screen bg-gray-50 flex items-center justify-center px-4 ">
      <form 
        onSubmit={handleSubmit}
        className="bg-white shadow-md border border-gray-200 rounded-xl w-full max-w-2xl p-8 space-y-6 my-20"
      >
        <h1 className="text-2xl font-semibold text-center text-gray-900">
          Register New SSO Auth Client
        </h1>

        {/* Admin-typed fields */}
        {["name", "client_id","client_secret","description", "auth_redirect_callback_uri", "sso_redirect_callback_uri", "scope"].map(field => (
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

      </form>
    </div>
  );
};

export default RegisterAuthClient;
