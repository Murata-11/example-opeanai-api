import { useState } from "react";
import axios from 'axios';
import { GetServerSideProps } from "next";
import { setCookie } from 'nookies';

export default function Home() {
  const [topic, setTopic] = useState("");
  const [ideas, setIdeas] = useState([]);
  const [loading, setLoading] = useState(false);

  const fetchIdeas = async (event: React.FormEvent) => {
    event.preventDefault();
    setLoading(true);
    setIdeas([]);
    try {
      const res = await axios.post("http://localhost:1323/brainstorm", { topic }, {
        headers: {
          "Content-Type": "application/json",
        },
        withCredentials: true,
      });
      const data = await res.data;
      setIdeas(data.ideas.slice(0, 5));
    } catch (error) {
      if (axios.isAxiosError(error) && error.response?.status === 403) {
        try {
          if (axios.isAxiosError(error) && error.response?.status === 403) {
            const refreshResponse = await axios.get("http://localhost:1323/refresh", {
              headers: {
                "Content-Type": "application/json",
              },
              withCredentials: true,
            });

            const newRes = await axios.post("http://localhost:1323/brainstorm", { topic }, {
              headers: {
                "Content-Type": "application/json",
              },
              withCredentials: true,
            });
            const data = await newRes.data;
            setIdeas(data.ideas.slice(0, 5));
          }
        } catch (refreshError) {
          if (axios.isAxiosError(refreshError) && refreshError.response?.status === 403) {
            return {
              redirect: {
                destination: "/login",
                permanent: false,
              },
            };
          }
          console.error("Error fetching ideas:", error);
        }
      }
    }
    setLoading(false);
  };

  return (
    <div className="min-h-screen flex flex-col items-center justify-center p-4 w-full">
      <h1 className="text-2xl font-bold mb-4">Brainstorming AI</h1>
      <input
        type="text"
        className="border p-2 rounded w-80 text-black"
        placeholder="Enter a topic..."
        value={topic}
        onChange={(e) => setTopic(e.target.value)}
      />
      <button
        onClick={fetchIdeas}
        className="bg-blue-500 text-white px-4 py-2 rounded mt-2"
        disabled={loading || !topic}
      >
        {loading ? "Generating..." : "Get Ideas"}
      </button>
      <div className="mt-4 grid grid-cols-1 sm:grid-cols-2 gap-4 w-full px-4">
        {ideas.map((idea, index) => (
          <div key={index} className="border p-4 rounded shadow-md bg-white w-full sm:w-10/12 lg:w-4/5 mx-auto">
            <p className="text-lg font-semibold text-black">{idea}</p>
          </div>
        ))}
      </div>
    </div>
  );
}

export const getServerSideProps: GetServerSideProps = async (context) => {
  const { req, res } = context;
  let cookie = req.headers.cookie;

  try {
    // 認可リクエストを送信
    const response = await axios.get("http://backend:1323/authz", {
      headers: {
        "Content-Type": "application/json",
        Cookie: cookie,
      },
    });

    return {
      props: {}, // will be passed to the page component as props
    }
  } catch (error) {
    if (axios.isAxiosError(error) && error.response?.status === 403) {
      try {
        const refreshResponse = await axios.get("http://backend:1323/refresh-server", {
          headers: {
            "Content-Type": "application/json",
            Cookie: cookie,
          },
          // withCredentials: true,
        });
        const newAccessToken = refreshResponse.data.accessToken;

        context.res.setHeader("Set-Cookie", `AccessToken=${newAccessToken}`);

        const retryResponse = await axios.get("http://backend:1323/authz", {
          headers: {
            "Content-Type": "application/json",
            Cookie: "AccessToken=" + newAccessToken,
          },
        });

        return {
          props: {}, // will be passed to the page component as props
        }
      } catch (refreshError) {
        if (axios.isAxiosError(refreshError) && refreshError.response?.status === 403) {
          return {
            redirect: {
              destination: "/login",
              permanent: false,
            },
          };
        }
      }
    }
  }

  return {
    props: {}, // will be passed to the page component as props
  };
};