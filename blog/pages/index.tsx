import { Container, Heading, Stack, Text } from "@chakra-ui/react";
import fs from "fs";
import matter from "gray-matter";
import Link from "next/link";
import Layout from "../components/Layout";

export default function Home({ posts }: any) {
  return (
    <Layout>
      <Container centerContent>
        <Stack spacing={8}>
          {posts.map(
            ({ frontmatter: { title, description, updatedAt }, slug }: any) => (
              <article key={slug}>
                <header>
                  <Heading style={{ cursor: "pointer" }}>
                    <Link href={"/posts/[slug]"} as={`/posts/${slug}`}>
                      <Text as="u">{title}</Text>
                    </Link>
                  </Heading>
                  <Text as="em" size="sm">
                    {updatedAt}
                  </Text>
                </header>
                <section>
                  <Text>{description}</Text>
                </section>
              </article>
            )
          )}
        </Stack>
      </Container>
    </Layout>
  );
}

export async function getStaticProps() {
  const files = fs.readdirSync(`${process.cwd()}/content/posts`);

  const posts = files.map((filename) => {
    const markdownWithMetadata = fs
      .readFileSync(`content/posts/${filename}`)
      .toString();

    const { data } = matter(markdownWithMetadata);

    // Convert post date to format: Month day, Year
    const options = { year: "numeric", month: "long", day: "numeric" };
    const formattedDate = data.updatedAt.toLocaleDateString("en-US", options);
    const frontmatter = {
      ...data,
      updatedAt: formattedDate,
    };

    return {
      slug: filename.replace(".md", ""),
      frontmatter,
    };
  });

  posts.sort((a, b) => {
    if (a.frontmatter.updatedAt > b.frontmatter.updatedAt) {
      return 1;
    } else {
      return -1;
    }
  });

  return {
    props: {
      posts,
    },
  };
}
