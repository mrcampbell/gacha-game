import { ChakraProvider, Container, extendTheme, Box } from "@chakra-ui/react";
import { mode } from "@chakra-ui/theme-tools";
import { Global } from "@emotion/react";
import React from "react";
import fontFamily from "../styles/font-face";

// 2. Extend the theme to include custom colors, fonts, etc
const colors = {
  brand: {
    900: "#1a365d",
    800: "#153e75",
    700: "#2a69ac",
  },
};

const styles = {
  global: (props: any) => ({
    body: {
      color: mode("gray.800", "whiteAlpha.900")(props),
      bg: mode("white", "gray.800")(props),
      lineHeight: "base",
    },
  }),
};

const fonts = {
  body: "'Inter', sans-serif",
  heading: "'Inter', sans-serif",
  mono: "'Inter', monospace",
};

const theme = extendTheme({ colors, styles, fonts });

export default function MyApp({ Component, pageProps }: any) {
  return (
    <ChakraProvider theme={theme}>
      <Global styles={fontFamily} />
      <Box padding="10">
        <Component {...pageProps} />
      </Box>
    </ChakraProvider>
  );
}
