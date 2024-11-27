"use client";

import { ErrorMessage } from "@/components/ErrorMessage";
import { PropsWithChildren, useActionState } from "react";
import { startTransition } from "react";

import { checkoutAction } from "../../actions";

export async function getCardHash({ cardName, cardNumber, expireDate, cvv }) {
  return Math.random().toString(36).substring(7);
}

export type CheckoutFormProps = {
  className?: string;
};

export function CheckoutForm(props: PropsWithChildren<CheckoutFormProps>) {

  const [state, formAction] = useActionState(checkoutAction, {
    error: null as string | null,
  });

  return (
    <form
      action={async (formData: FormData) => {
        const card_hash = await getCardHash({
          cardName: formData.get("card_name") as string,
          cardNumber: formData.get("cc") as string,
          expireDate: formData.get("expire_date") as string,
          cvv: formData.get("cvv") as string,
        });
        startTransition(() => {
          formAction({
            cardHash: card_hash,
            email: formData.get("email") as string,
          });
        });
      }}
      className={props.className}
    >
      {state?.error && <ErrorMessage error={state.error} />}
      <input type="hidden" name="card_hash" />
      {props.children}
    </form>
  );
}