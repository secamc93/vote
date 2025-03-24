// src/config/env.ts

let _env: { apiUrl?: string };

export function getRuntimeEnv() {
  if (_env) return _env;

  const apiUrl = process.env.NEXT_PUBLIC_URL_API;

  if (!apiUrl || apiUrl.trim() === "") {
    throw new Error(
      `❌ La variable de entorno "NEXT_PUBLIC_URL_API" no está definida o está vacía. ` +
      `Asegúrate de que existe en tu archivo .env.local en la raíz del proyecto.`
    );
  }

  _env = { apiUrl };
  return _env;
}
