import {
	useEffect,
	useState,
} from "react";

import {
	useParams,
} from "react-router-dom";

import api from "../services/api";

function FileDetails() {

	const { fileName } =
		useParams();

	const [file,
		setFile] =
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
						`/api/file/${fileName}`,
						{
							headers: {
								Authorization:
									`Bearer ${token}`,
							},
						},
					);

				setFile(
					response.data,
				);
			};

		fetchData();

	}, [fileName]);

	if (!file) {

		return <h2>
			Loading...
		</h2>;
	}

	return (

		<div
			className="card"
		>

			<h1>
				File Details
			</h1>

			<p>
				<b>File:</b>
				{" "}
				{file.file_name}
			</p>

			<p>
				<b>SHA256:</b>
				{" "}
				{file.sha256}
			</p>

			<p>
				<b>Risk Score:</b>
				{" "}
				{file.risk_score}
			</p>

			<p>
				<b>Risk Level:</b>
				{" "}
				{file.risk_level}
			</p>

			<p>
				<b>Verdict:</b>
				{" "}
				{file.verdict}
			</p>

			<h3>
				MITRE ATT&CK
			</h3>

			<pre>
				{file.mitre}
			</pre>

			<h3>
				Findings
			</h3>

			<pre>
				{file.findings}
			</pre>

			<h3>
				URLs
			</h3>

			<pre>
				{file.urls}
			</pre>

		</div>
	);
}

export default FileDetails;