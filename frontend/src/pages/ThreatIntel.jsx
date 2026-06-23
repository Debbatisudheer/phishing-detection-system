import { useEffect, useState } from "react";
import api from "../services/api";

function ThreatIntel() {

const [stats, setStats] =
useState({});

const [files, setFiles] =
useState([]);

const [iocs, setIOCs] =
useState([]);

const [iocInput, setIOCInput] =
useState("");

const [reputation, setReputation] =
useState(null);

useEffect(() => {

const fetchData = async () => {

  try {

    const token =
      localStorage.getItem(
        "token"
      );

    const response =
      await api.get(
        "/api/threat-intel",
        {
          headers: {
            Authorization:
              "Bearer " + token,
          },
        }
      );

    setStats(
      response.data.stats
    );

    setFiles(
      response.data.top_files
    );

    setIOCs(
      response.data.top_iocs
    );

  } catch (error) {

    console.error(
      error
    );

  }
};

fetchData();

}, []);

const lookupIOC =
async () => {

  try {

    const response =
      await api.get(
        "/api/ioc-reputation?ioc=" +
          iocInput
      );

    setReputation(
      response.data
    );

  } catch (error) {

    console.error(
      error
    );

  }
};

return (

<div
  style={{
    padding: "20px",
  }}
>

  <h1>
    Threat Intelligence Dashboard
  </h1>

  <div
    style={{
      display: "grid",
      gridTemplateColumns:
        "repeat(3,1fr)",
      gap: "20px",
      marginBottom: "30px",
    }}
  >

    <div className="card">
      <h3>Total IOCs</h3>
      <h1>
        {stats.total_iocs}
      </h1>
    </div>

    <div className="card">
      <h3>Total Alerts</h3>
      <h1>
        {stats.total_alerts}
      </h1>
    </div>

    <div className="card">
      <h3>Critical Files</h3>
      <h1>
        {stats.critical_files}
      </h1>
    </div>

  </div>

  <button
    style={{
      padding: "12px 20px",
      marginBottom: "25px",
      cursor: "pointer",
    }}
    onClick={() =>
      window.open(
        "http://localhost:8081/api/export-report",
        "_blank"
      )
    }
  >
    Download PDF Report
  </button>

  <h2>
    Top Risk Files
  </h2>

  {files.map(
    (
      item,
      index
    ) => (

      <div
        key={index}
        className="card"
        style={{
          borderLeft:
            item.level ===
            "CRITICAL"
              ? "5px solid red"
              : "5px solid orange",
          marginBottom:
            "15px",
        }}
      >

        <h3>
          {item.file}
        </h3>

        <p>
          Risk:
          {" "}
          {item.level}
        </p>

        <p>
          Score:
          {" "}
          {item.score}
        </p>

      </div>

    )
  )}

  <h2>
    Top IOCs
  </h2>

  <table
    style={{
      width: "100%",
      borderCollapse:
        "collapse",
      marginBottom:
        "30px",
    }}
  >

    <thead>

      <tr>

        <th
          style={{
            border:
              "1px solid gray",
            padding:
              "10px",
          }}
        >
          IOC
        </th>

        <th
          style={{
            border:
              "1px solid gray",
            padding:
              "10px",
          }}
        >
          Hits
        </th>

      </tr>

    </thead>

    <tbody>

      {iocs.map(
        (
          item,
          index
        ) => (

          <tr
            key={index}
          >

            <td
              style={{
                border:
                  "1px solid gray",
                padding:
                  "10px",
              }}
            >
              {item.ioc}
            </td>

            <td
              style={{
                border:
                  "1px solid gray",
                padding:
                  "10px",
              }}
            >
              {item.count}
            </td>

          </tr>

        )
      )}

    </tbody>

  </table>

  <h2>
    IOC Reputation Lookup
  </h2>

  <input
    type="text"
    value={iocInput}
    onChange={
      (e) =>
        setIOCInput(
          e.target.value
        )
    }
    placeholder="Enter IOC"
    style={{
      padding: "10px",
      marginRight:
        "10px",
      width: "300px",
    }}
  />

  <button
    onClick={
      lookupIOC
    }
  >
    Lookup
  </button>

  {reputation && (

    <div
      className="card"
      style={{
        marginTop:
          "20px",
      }}
    >

      <p>
        Reputation:
        {" "}
        {reputation.reputation}
      </p>

      <p>
        Source:
        {" "}
        {reputation.source}
      </p>

    </div>

  )}

</div>

);
}

export default ThreatIntel;
