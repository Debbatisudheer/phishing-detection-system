import { useState } from "react";
import api from "../services/api";
import "./Search.css";

function Search() {
  const [query, setQuery] = useState("");
  const [results, setResults] = useState([]);

  const searchIOC = async () => {
    const token = localStorage.getItem("token");

    const response = await api.get(`/api/search?q=${query}`, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });

    setResults(
      Array.isArray(response.data)
        ? response.data
        : [],
    );
  };

  return (
    <main className="search-page">
      <section className="search-header">
        <h1>IOC Search</h1>
      </section>

      <section className="search-panel">
        <div className="search-controls">
          <input
            placeholder="Search IOC"
            value={query}
            onChange={(e) => setQuery(e.target.value)}
          />

          <button onClick={searchIOC}>
            Search
          </button>
        </div>
      </section>

      {results.length > 0 ? (
        <section className="search-results">
          {results.map((r, index) => (
            <article
              key={index}
              className="search-result-card"
            >
              <div className="result-card-header">
                <h3>{r.file_name}</h3>
              </div>

              <div className="result-details">
                <div className="result-detail">
                  <span>Risk</span>
                  <strong
                    className={`risk-${String(
                      r.risk_level || "",
                    ).toLowerCase()}`}
                  >
                    {r.risk_level}
                  </strong>
                </div>

                <div className="result-detail">
                  <span>Verdict</span>
                  <strong>{r.verdict}</strong>
                </div>

                <div className="result-detail result-detail-wide">
                  <span>SHA256</span>
                  <strong className="hash-value">
                    {r.sha256}
                  </strong>
                </div>

                <div className="result-detail">
                  <span>MITRE</span>
                  <strong>{r.mitre}</strong>
                </div>
              </div>
            </article>
          ))}
        </section>
      ) : (
        query !== "" && (
          <section className="no-results">
            <div className="no-results-icon">⌕</div>
            <h3>No results found</h3>
          </section>
        )
      )}
    </main>
  );
}

export default Search;