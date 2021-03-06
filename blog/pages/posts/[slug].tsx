import fs from "fs";
import path from "path";
import matter from "gray-matter";
import ReactMarkdown from "react-markdown/with-html";
import { Prism as SyntaxHighlighter } from "react-syntax-highlighter";
import theme from "react-syntax-highlighter/dist/cjs/styles/prism/atom-dark";
import { Container, Box } from "@chakra-ui/react";

// @ts-ignore
import ChakraUIRenderer from "chakra-ui-markdown-renderer";

import Layout from "../../components/Layout";

const CodeBlock = ({ language, value }: any) => {
  return (
    <Box maxW={10}>
      <SyntaxHighlighter language={language} style={theme}>
        {value}
      </SyntaxHighlighter>
    </Box>
  );
};

export default function Post({ content, frontmatter }: any) {
  return (
    <Layout>
      <article style={{padding: 40}}>
          <ReactMarkdown
            escapeHtml={false}
            source={content}
            renderers={{ code: CodeBlock, ...ChakraUIRenderer() }}
          />
      </article>
    </Layout>
  );
}

export async function getStaticPaths() {
  const files = fs.readdirSync("content/posts");

  const paths = files.map((filename) => ({
    params: {
      slug: filename.replace(".md", ""),
    },
  }));

  return {
    paths,
    fallback: false,
  };
}

export async function getStaticProps({ params: { slug } }: any) {
  const markdownWithMetadata = fs
    .readFileSync(path.join("content/posts", slug + ".md"))
    .toString();

  const { data, content } = matter(markdownWithMetadata);

  // Convert post date to format: Month day, Year
  const options = { year: "numeric", month: "long", day: "numeric" };
  const formattedDate = data.updatedAt.toLocaleDateString("en-US", options);

  const frontmatter = {
    ...data,
    updatedAt: formattedDate,
  };

  return {
    props: {
      content,
      frontmatter,
    },
  };
}
