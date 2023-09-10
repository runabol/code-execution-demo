"use client";

import { Editor } from "@monaco-editor/react";
import { useState } from "react";

export default function Home() {
  const [code, setCode] = useState("");
  const [language, setLanguage] = useState("bash");
  const [output, setOutput] = useState("Output:");
  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24">
      <div className="z-10 max-w-5xl w-full items-center justify-between font-mono text-sm lg:flex">
        <div className="w-full flex flex-col gap-2">
          <div className="flex">
            <div className="flex w-full justify-end  align-middle gap-2">
              <select
                id="location"
                name="location"
                className="block rounded-md border-0 py-1.5 pl-3 pr-10 text-gray-900 ring-1 ring-inset ring-gray-300 focus:ring-2 focus:ring-indigo-600 sm:text-sm sm:leading-6"
                defaultValue={language}
                onChange={(ev) => setLanguage(ev.target.value)}
              >
                <option value="bash">Bash</option>
                <option value="python3">Python</option>
                <option value="golang">Golang</option>
              </select>
              <button
                type="button"
                className="rounded bg-indigo-600 px-3 py-2 text-md font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600 "
                onClick={() => {
                  fetch("/api/execute", {
                    method: "POST",
                    body: JSON.stringify({
                      langauge: language,
                      code: code,
                    }),
                    headers: {
                      "content-type": "application/json",
                    },
                  })
                    .then((res) => res.json())
                    .then((data) => {
                      if (data.message) {
                        setOutput(data.message);
                      } else {
                        setOutput(data.output);
                      }
                    });
                }}
              >
                Run
              </button>
            </div>
          </div>
          <Editor
            height="25rem"
            language="bash"
            theme="vs-dark"
            className="max-w-5xl"
            value={code}
            onChange={(value) => setCode(value || "")}
          />
          <div className="overflow-hidden bg-white shadow sm:rounded-lg">
            <div className="px-4 py-5 sm:p-6">{output}</div>
          </div>
        </div>
      </div>
    </main>
  );
}
