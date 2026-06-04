import { useEffect, useState } from "react";
import api from "../services/api";

function Cases() {

  const [cases, setCases] =
    useState([]);

  useEffect(() => {

    const fetchCases =
      async () => {

        const token =
          localStorage.getItem(
            "token",
          );

        const response =
          await api.get(
            "/api/cases",
            {
              headers: {
                Authorization:
                  `Bearer ${token}`,
              },
            },
          );

        setCases(
          response.data,
        );
      };

    fetchCases();

  }, []);

  return (
    <div>

      <h1>Cases</h1>

      {cases.map(
        (c) => (

          <div
            key={c.id}
            style={{
              border:
                "1px solid black",
              margin: "10px",
              padding: "10px",
            }}
          >

            <h3>
              {c.file_name}
            </h3>

            <p>
              Analyst:
              {c.analyst}
            </p>

            <p>
              Status:
              {c.status}
            </p>

            <p>
              Notes:
              {c.notes}
            </p>

          </div>
        ),
      )}

    </div>
  );
}

export default Cases;