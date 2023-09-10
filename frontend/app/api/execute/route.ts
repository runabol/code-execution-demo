import { NextResponse } from "next/server";

export async function POST(request: Request) {
  const body = await request.text();
  const res = await fetch(`${process.env.BACKEND_URL}/execute`, {
    method: "POST",
    body: body,
    headers: {
      "content-type": "application/json",
    },
  });
  const data = await res.json();
  if (res.ok) {
    return NextResponse.json(data);
  } else {
    return NextResponse.json(data, { status: res.status });
  }
}
