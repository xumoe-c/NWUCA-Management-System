import http from './http';

export interface Department {
    id: number;
    name: string;
    createdAt?: string;
    updatedAt?: string;
}

export async function getDepartments() {
    const res = await http.get<Department[]>('/departments');
    return res.data;
}

export async function createDepartment(payload: { name: string }) {
    const res = await http.post<Department>('/departments/', payload);
    return res.data;
}

export async function deleteDepartment(id: number) {
    const res = await http.delete(`/departments/${id}`);
    return res.data;
}
