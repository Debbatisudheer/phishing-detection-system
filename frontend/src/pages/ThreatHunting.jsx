import { useEffect, useState }
from "react";

import api
from "../services/api";

function ThreatHunting() {

  const [data, setData] =
    useState(null);

  useEffect(() => {

    const fetchData =
      async () => {

        const token =
          localStorage.getItem(
            "token",
          );

        const response =
          await api.get(
            "/api/threat-hunting",
            {
              headers: {
                Authorization:
                  `Bearer ${token}`,
              },
            },
          );

        setData(
          response.data,
        );
      };

    fetchData();

  }, []);

  if (!data)
    return <h2>Loading...</h2>;

  return (

    <div>

      <h1>
        Threat Hunting
      </h1>

      <h3>
        Critical Files:
        {data.critical_files}
      </h3>

      <h3>
        Quarantine Files:
        {data.quarantine_files}
      </h3>

      <h2>
        MITRE Techniques
      </h2>

      {Object.entries(
        data.top_mitre,
      ).map(
        ([technique, count]) => (

          <p
            key={technique}
          >

            {technique}
            {" : "}
            {count}

          </p>
        ),
      )}

    </div>
  );
}

export default ThreatHunting;