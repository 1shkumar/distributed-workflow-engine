CREATE TABLE workflows (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    definition JSONB NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE workflow_runs (
    id UUID PRIMARY KEY,
    workflow_id UUID REFERENCES workflows(id),
    status TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE task_runs (
    id UUID PRIMARY KEY,
    workflow_run_id UUID REFERENCES workflow_runs(id),
    task_name TEXT NOT NULL,
    status TEXT NOT NULL,
    attempt INT DEFAULT 0,
    started_at TIMESTAMP,
    completed_at TIMESTAMP
);