// SSO API Configuration
const API_BASE_URL = import.meta.env.VITE_SSO_API_URL || 'http://localhost:8080';

export interface AuthorizeRequest {
  client_id: string;
  redirect_uri: string;
  response_type: string;
  scope?: string;
  state?: string;
}

export interface LoginRequest {
  username: string;
  password: string;
  client_id?: string;
  redirect_uri?: string;
  csrf_ses?: string;
  rid?: string;
}

export interface TokenRequest {
  grant_type: 'authorization_code';
  code: string;
  redirect_uri: string;
  client_id: string;
  client_secret: string;
}

export interface LoginResponse {
  redirect_uri: string;
  code?: string;
  error?: string;
  nonce?: string;
}

export interface TokenResponse {
  access_token: string;
  token_type: 'Bearer';
  expires_in: number;
  refresh_token?: string;
  scope?: string;
}

// API Functions
export const ssoApi = {
  // Validate authorization request
  // authorize: async (params: AuthorizeRequest): Promise<{ valid: boolean; error?: string }> => {
  //   try {
  //     const response = await fetch(`${API_BASE_URL}/authorize`, {
  //       method: 'POST',
  //       headers: { 'Content-Type': 'application/json' },
  //       body: JSON.stringify(params),
  //     });
  //     return await response.json();
  //   } catch (error) {
  //     console.error('Authorize API error:', error);
  //     return { valid: false, error: 'Network error' };
  //   }
  // },

  // Authenticate user and get authorization code
  login: async (params: LoginRequest): Promise<LoginResponse> => {
    try {
      const response = await fetch(`/api/sso/login`, {
        method: 'POST',
        headers: { 
          'Content-Type': 'application/json',
          'X-csrf_token': params.csrf_ses || '',
        },
        body: JSON.stringify(params),
      });
      return await response.json();
    } catch (error) {
      console.error('Login API error:', error);
      return {redirect_uri: '', error: 'Network error' };
    }
  },

  // Exchange authorization code for access token
  token: async (params: TokenRequest): Promise<TokenResponse> => {
    try {
      const response = await fetch(`$/aoi/sso/token`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(params),
      });
      return await response.json();
    } catch (error) {
      console.error('Token API error:', error);
      throw new Error('Failed to exchange code for token');
    }
  },

  // Get user info with access token
  // userInfo: async (accessToken: string): Promise<any> => {
  //   try {
  //     const response = await fetch(`${API_BASE_URL}/userinfo`, {
  //       headers: { 'Authorization': `Bearer ${accessToken}` },
  //     });
  //     return await response.json();
  //   } catch (error) {
  //     console.error('UserInfo API error:', error);
  //     throw new Error('Failed to get user info');
  //   }
  // }
};
