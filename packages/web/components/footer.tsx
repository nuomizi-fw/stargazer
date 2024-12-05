import { Link } from "@nextui-org/link";

export const Footer = () => {
  return (
    <footer className="w-full flex items-center justify-center py-3">
      <Link
        isExternal
        className="flex items-center gap-1 text-current"
        href="https://nextui-docs-v2.vercel.app?utm_source=next-app-template"
        title="Stargazer"
      >
        <span className="text-default-600">Powered by</span>
        <p className="text-primary">NextUI</p>
      </Link>
    </footer>
  );
};
