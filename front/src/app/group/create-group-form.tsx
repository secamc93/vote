'use client';

import { useState } from "react";
import { createVoteGroupUseCases } from "@/application/usecase/constructor";
import QRCode from "react-qr-code";
import styles from "./page.group.module.css";

export default function CreateGroupForm() {
  const [name, setName] = useState("");
  const [message, setMessage] = useState("");
  const [qr, setQr] = useState<string | null>(null);
  
  const voteGroupUseCases = createVoteGroupUseCases(); // Nueva instancia

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      const response = await voteGroupUseCases.createVoteGroup({ name }); // Se usa la instancia
      setMessage(`✅ ${response.message} (ID: ${response.group_id})`);
      setQr(response.group_id.toString()); // seteamos el id para el QR
      setName(""); // limpia el campo
    } catch (error: unknown) {
      if (error instanceof Error) {
        setMessage(`❌ ${error.message}`);
      } else {
        setMessage("❌ An unexpected error occurred.");
      }
      setQr(null);
    }
  };

  return (
    <div className={styles.form}>
      {qr ? (
        <div className={styles.qrContainer}>
          <QRCode value={qr} />
        </div>
      ) : (
        <>
          <form onSubmit={handleSubmit}>
            <label>
              Nombre del grupo
              <input
                type="text"
                value={name}
                onChange={(e) => setName(e.target.value)}
                required
              />
            </label>
            <button type="submit">Crear grupo</button>
          </form>
          {message && <p>{message}</p>}
        </>
      )}
    </div>
  );
}
