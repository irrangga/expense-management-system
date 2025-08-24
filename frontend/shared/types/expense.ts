export interface Expense {
  id: number;
  user_id: number;
  amount_idr: number;
  description: string;
  receipt_url: string;
  status: string;
  submitted_at: string;
  processed_at: string;
}
