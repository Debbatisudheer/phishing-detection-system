import { useEffect, useState } from "react";
import api from "../services/api";

function IOCGraph() {

  const [data, setData] =
    useState([]);

  useEffect(() => {

    loadGraph();

  }, []);

  const loadGraph =
    async () => {

      const token =
        localStorage.getItem(
          "token",
        );

      const response =
        await api.get(
          "/api/ioc-graph",
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

  return (

    <div>

      <h1>
        IOC Relationship View
      </h1>

      <table
        border="1"
        width="100%"
      >

        <thead>

          <tr>

            <th>IOC</th>

            <th>Source</th>

            <th>File</th>

          </tr>

        </thead>

        <tbody>

          {data.map(
            (item, index) => (

              <tr key={index}>

                <td>
                  {item.ioc}
                </td>

                <td>
                  {item.source}
                </td>

                <td>
                  {item.file}
                </td>

              </tr>
            ),
          )}

        </tbody>

      </table>

    </div>
  );
}

export default IOCGraph;