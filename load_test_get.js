import http from 'k6/http';
import { check, sleep } from 'k6';
import { SharedArray } from 'k6/data';

// 1. Load the IDs into a SharedArray (memory efficient)
const ids = new SharedArray('customer ids', function () {
    return [
        'ffffc8bb-a22b-414f-93c5-aabd02545f4e',
        'ffc8d561-e2f8-40b3-9fbf-81297501d59e',
        'ffbcd58e-04c7-4145-b056-e72fd6f431b0',
        'ffa3edb7-1645-4c76-90ed-0a01a92e8aeb',
        'ff87c278-7621-4451-ad84-af9577338363',
        'ff57fc66-2c51-40f4-ae2c-9e01c6498a19',
        'ff54172f-8478-4e15-af53-4a82f23ab7b3',
        'ff35f8d2-8f24-4944-8fd5-5ff51ec05f9c',
        'ff27eaa0-34ce-4476-9235-1e964571156b',
        'fe852a34-ec0f-41c6-b4cd-8577d098af8b',
    ];
});

export const options = {
    // 2. Define the load pattern
    stages: [
        { duration: '30s', target: 20 }, // Ramp up to 20 users
        { duration: '1m', target: 20 },  // Stay at 20 users
        { duration: '10s', target: 0 },  // Ramp down
    ],
    thresholds: {
        http_req_failed: ['rate<0.01'], // Fail if more than 1% errors
        http_req_duration: ['p(95)<200'], // 95% of requests should be < 200ms
    },
};

export default function () {
    // 3. Select a random ID from the list
    const randomId = ids[Math.floor(Math.random() * ids.length)];
    const url = `http://localhost:8080/go-hexagonal-example/v1/customer-profile/${randomId}`;

    const params = {
        headers: {
            'Content-Type': 'application/json',
        },
    };

    const res = http.get(url, params);

    // 4. Validate the response
    check(res, {
        'status is 200': (r) => r.status === 200,
        'has correct body': (r) => r.body.includes(randomId),
    });

    // Think time between requests (simulate human behavior)
    sleep(1);
}