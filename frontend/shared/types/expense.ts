export interface Expense {
  id: number;
  user_name: string;
  amount_idr: number;
  description: string;
  receipt_url: string;
  status: string;
  requires_approval: boolean;
  auto_approved: boolean;
  submitted_at: string;
  processed_at: string;
}
