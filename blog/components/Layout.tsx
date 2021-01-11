import Link from "next/link";
import { useRouter } from "next/router";
import { Heading, Box } from "@chakra-ui/react";

export default function Layout({ children }: any) {
  const { pathname } = useRouter();
  const isRoot = pathname === "/";

  const header = isRoot ? (
    <Box marginBottom={8}>
      <Link href="/">
        <Heading>
          MC's Blog
        </Heading>
      </Link>
    </Box>
  ) : (
    <Box marginBottom={2}>
      <Link href="/">
        <Heading>
          MC's Blog
        </Heading>
      </Link>
    </Box>
  );

  return (
    <Box w="100%">
      <header>{header}</header>
      <Box w="80%">{children}</Box>
      <footer>
        <Box marginTop={12}>
        © {new Date().getFullYear()}, Built with{" "}
        <a href="https://nextjs.org/">Next.js</a> &#128293;
        </Box>
      </footer>
    </Box>
  );
}