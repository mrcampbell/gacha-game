import React, { useState } from "react";
import { Grid, GridItem } from "@chakra-ui/react";
import { AnimatePresence, AnimateSharedLayout, motion } from "framer-motion";

export const MainMenu = () => {
  const [selectedId, setSelectedId] = useState(null);

  return (
    <AnimateSharedLayout type="crossfade">
      <Grid
        h="200px"
        minWidth="100%"
        templateRows="repeat(2, 1fr)"
        templateColumns="repeat(5, 1fr)"
        gap={4}>
        <GridItem rowSpan={2} colSpan={1} bg="tomato" />
        <GridItem colSpan={2} bg="papayawhip" />
        <GridItem colSpan={2} bg="papayawhip" />
        <GridItem colSpan={4} bg="tomato" />
      </Grid>
      <AnimatePresence>
        {selectedId && (
          <motion.div layoutId="underline">
            <motion.h5>"A"</motion.h5>
            <motion.h2>"B"</motion.h2>
            <motion.button onClick={() => setSelectedId(null)} />
          </motion.div>
        )}
      </AnimatePresence>
    </AnimateSharedLayout>
  );
};
