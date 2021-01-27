import * as React from "react";
import {
  ChakraProvider,
  Box,
  VStack,
  theme,
  Heading,
} from "@chakra-ui/react";
import { ColorModeSwitcher } from "./ColorModeSwitcher";
import {MainMenu} from "../components/MainMenu/MainMenu"

export const App = () => (
  <ChakraProvider theme={theme}>
    <VStack fontSize="xl">
      <Box
        display="flex"
        flexDir="row"
        justifyContent="space-between"
        minW="100vw"
        p="3"
        border="1px solid purple">
        <Heading>Overkill: The Game</Heading>
        <ColorModeSwitcher />
      </Box>
      <Box className="body" minW="100vw" p="3"> 

      <MainMenu />
      </Box>
    </VStack>
  </ChakraProvider>
);
