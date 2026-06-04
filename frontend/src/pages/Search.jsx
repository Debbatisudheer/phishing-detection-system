import { useState } from "react";
import api from "../services/api";

function Search() {

  const [query, setQuery] =
    useState("");

  const [results, setResults] =
    useState([]);

  const searchIOC =
    async () => {

      const token =
        localStorage.getItem(
          "token",
        );

      const response =
        await api.get(
          `/api/search?q=${query}`,
          {
            headers: {
              Authorization:
                `Bearer ${token}`,
            },
          },
        );

      setResults(
        response.data,
      );
    };

  return (
    <div>

      <h1>
        IOC Search
      </h1>

      <input
        placeholder="Search IOC"
        value={query}
        onChange={(e) =>
          setQuery(
            e.target.value,
          )
        }
      />

      <button
        onClick={searchIOC}
      >
        Search
      </button>

      {results.map(
        (r, index) => (

          <div
            key={index}
            style={{
              border:
                "1px solid black",
              margin: "10px",
              padding: "10px",
            }}
          >

            <h3>
              {r.file_name}
            </h3>

            <p>
              Risk:
              {r.risk_level}
            </p>

            <p>
              Verdict:
              {r.verdict}
            </p>

            <p>
              SHA256:
              {r.sha256}
            </p>

            <p>
              MITRE:
              {r.mitre}
            </p>

          </div>
        ),
      )}

    </div>
  );
}

export default Search;