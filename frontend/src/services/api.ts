const API_BASE = '/api';

interface ErrorBody {
  code?: string;
  message?: string;
}

// request 统一处理 API 请求，后端业务错误会抛出带 message 的 Error。
export async function request<T>(path: string, options?: RequestInit): Promise<T> {
  const response = await fetch(`${API_BASE}${path}`, {
    headers: { 'Content-Type': 'application/json' },
    ...options,
  });
  if (!response.ok) {
    let message = '请求失败，请稍后重试';
    try {
      const body = (await response.json()) as ErrorBody;
      if (body.message) {
        message = body.message;
      }
    } catch {
      // 非 JSON 错误体时保留默认提示
    }
    throw new Error(message);
  }
  return response.json() as Promise<T>;
}
