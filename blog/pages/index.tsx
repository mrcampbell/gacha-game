import fs from "fs";
import matter from "gray-matter";
import Link from "next/link";
import Layout from "../components/Layout";

export default function Home({ posts }: any) {
  return (
    <Layout>
      {posts.map(({ frontmatter: { title, description, date }, slug }: any) => (
        <article key={slug}>
          <header>
            <h3 className="mb-2">
              <Link href={"/posts/[slug]"} as={`/posts/${slug}`}>
                <a className="text-3xl font-semibold text-orange-600 no-underline">
                  {title}
                </a>
              </Link>
            </h3>
            <span className="mb-4 text-xs">{date}</span>
          </header>
          <section>
            <p className="mb-8">{description}</p>
          </section>
        </article>
      ))}
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
    console.log(data)
    const formattedDate = data.updatedAt.toLocaleDateString("en-US", options);
    console.log(formattedDate)
    const frontmatter = {
      ...data,
      updatedAt: formattedDate,
    };

    return {
      slug: filename.replace(".md", ""),
      frontmatter,
    };
  });

  return {
    props: {
      posts,
    },
  };
}