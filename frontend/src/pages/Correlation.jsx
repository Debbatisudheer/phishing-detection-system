import { useEffect, useState } from "react";
import api from "../services/api";

function Correlation() {

  const [data, setData] =
    useState([]);

  useEffect(() => {

    loadData();

  }, []);

  const loadData =
    async () => {

      const token =
        localStorage.getItem(
          "token",
        );

      const response =
        await api.get(
          "/api/correlation",
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
        IOC Correlation
      </h1>

      <table
        border="1"
        width="100%"
      >

        <thead>

          <tr>

            <th>IOC</th>

            <th>Occurrences</th>

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
                  {item.count}
                </td>

              </tr>
            ),
          )}

        </tbody>

      </table>

    </div>
  );
}

export default Correlation;