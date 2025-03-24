declare module "react-qr-code" {
	interface QRCodeProps {
		value: string;
		// ...otras props opcionales si son necesarias
	}
	const QRCode: React.FC<QRCodeProps>;
	export default QRCode;
}
