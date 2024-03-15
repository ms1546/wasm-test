use wasm_bindgen::prelude::*;

#[wasm_bindgen]
pub fn calculate(a: f64, b: f64, operation: &str) -> String {
    let result = match operation {
        "add" => a + b,
        "subtract" => a - b,
        "multiply" => a * b,
        "divide" => {
            if b == 0.0 {
                return "Cannot divide by zero!".to_string();
            } else {
                a / b
            }
        },
        _ => return "Invalid operation".to_string(),
    };
    format!("Result: {}", result)
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_addition() {
        assert_eq!(calculate(2.0, 3.0, "add"), "Result: 5");
    }

    #[test]
    fn test_subtraction() {
        assert_eq!(calculate(5.0, 3.0, "subtract"), "Result: 2");
    }

    #[test]
    fn test_multiplication() {
        assert_eq!(calculate(3.0, 4.0, "multiply"), "Result: 12");
    }

    #[test]
    fn test_division() {
        assert_eq!(calculate(8.0, 2.0, "divide"), "Result: 4");
    }

    #[test]
    fn test_divide_by_zero() {
        assert_eq!(calculate(1.0, 0.0, "divide"), "Cannot divide by zero!");
    }

    #[test]
    fn test_invalid_operation() {
        assert_eq!(calculate(1.0, 1.0, "invalid"), "Invalid operation");
    }
}
