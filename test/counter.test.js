import { check } from "k6";
import { get, increment, decrement, add, set, reset } from "k6/x/example";

export const options = {
  thresholds: {
    checks: ["rate==1"],
  },
};

export default function () {
  check(get(), {
    "counter is 0": (counter) => counter === 0,
  });

  increment()
  check(get(), {
    "counter is 1": (counter) => counter === 1,
  });

  add(9)
  check(get(), {
    "counter is 10": (counter) => counter === 10,
  });

  reset()
  check(get(), {
    "counter is 0": (counter) => counter === 0,
  });

  set(-99)
  check(get(), {
    "counter is -99": (counter) => counter === -99,
  });

  decrement()
  check(get(), {
    "counter is -100": (counter) => counter === -100,
  });
}
