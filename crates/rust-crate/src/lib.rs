pub fn add(a: i32, b: i32) -> i32 {
    a + b
}

pub fn generate_nanoid() -> String {
    nanoid::nanoid!()
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_add() {
        assert_eq!(add(1, 2), 3);
    }

    #[test]
    fn test_generate_nanoid() {
        assert_eq!(generate_nanoid().len(), 21);
    }
}
