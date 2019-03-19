export interface User {
    id: string
    created_at: string
    updated_at: string
    email: string
    github_id: string
    github_user_name: string
    first_name: string
    last_name: string
}

export interface Repo {
    id: string
    owner: string
    name: string
    description?: string
    url: string
}
