import type { Metadata } from 'next';
import './globals.css';

export const metadata: Metadata = {
  title: 'Next Ecommerce',
  description: 'Next Project Ecommerce',
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body>{children}</body>
    </html>
  );
}
