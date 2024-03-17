INSERT INTO restaurants (name, description, address, rating, restaurant_type, num_of_reviews, image_url)
VALUES
  ('Chicken Rice Express', 'Famous for Hainanese Chicken Rice', '123 Orchard Road, Singapore', 5.0, 'Chicken', 10000, 'chicken_rice_image_url'),
  ('Noodle Delight', 'Specializing in various noodle dishes', '456 Raffles Place, Singapore', 5.0, 'Chinese', 10000, 'noodle_image_url'),
  ('Sushi Haven', 'Authentic Japanese sushi and sashimi', '789 Marina Bay Sands, Singapore', 5.0, 'Japanese', 10000, 'sushi_image_url'),
  ('Spicy Noodle House', 'Home of spicy noodle lovers', '101 Tanjong Pagar Road, Singapore', 5.0, 'Korean', 10000, 'spicy_noodle_image_url'),
  ('Pizza Paradise', 'Delicious pizzas with a variety of toppings', '202 Sentosa Island, Singapore', 5.0, 'Pizza', 10000, 'pizza_image_url'),
  ('Curry House', 'Flavorful Indian curries', '303 Little India, Singapore', 5.0, 'Indian', 10000, 'curry_image_url'),
  ('Thai Delights', 'Traditional Thai dishes with a modern twist', '404 Chinatown, Singapore', 5.0, 'Thai', 10000, 'thai_image_url'),
  ('Dessert Dreamland', 'Satisfy your sweet tooth with our desserts', '505 Bugis Street, Singapore', 5.0, 'Dessert', 10000, 'dessert_image_url'),
  ('Healthy Bites', 'Nutritious and delicious healthy options', '606 Orchard Boulevard, Singapore', 5.0, 'Healthy', 10000, 'healthy_image_url'),
  ('Western Grill', 'Classic Western dishes for all tastes', '707 Clarke Quay, Singapore', 5.0, 'Western', 10000, 'western_image_url'),
  ('Satay Street', 'Famous for delicious satay skewers', '111 Marina Bay, Singapore', 5.0, 'Satay', 10000, 'satay_image_url'),
  ('Dim Sum Delight', 'Authentic Cantonese dim sum experience', '222 Chinatown, Singapore', 5.0, 'Chinese', 10000, 'dim_sum_image_url'),
  ('Ramen House', 'Specializing in various ramen varieties', '333 Orchard Road, Singapore', 5.0, 'Japanese', 10000, 'ramen_image_url');





INSERT INTO dishes (restaurant_id, is_available, name, description, price, diet_type, cuisine, image_url)
VALUES

  -- Chicken Rice Express
  (1, true, 'Hainanese Chicken Rice', 'Classic Hainanese chicken with fragrant rice', 5.50, 'normal', '{"chicken", "rice"}', 'chicken_rice_image_url_1'),
  (1, true, 'Roast Chicken Noodles', 'Roasted chicken served with noodles', 6.50, 'normal', 'chicken', 'chicken_rice_image_url_2'),
  (1, true, 'Chicken Satay Rice', 'Hainanese-style chicken rice with chicken satay', 8.50, 'normal', 'chicken','chicken_rice_image_url_3'),
  (1, true, 'Spicy Chicken Noodles', 'Spicy chicken noodles with a kick', 7.00, 'normal', 'chicken', 'chicken_rice_image_url_4'),

  -- Noodle Delight
  (2, true, 'Singapore-style Noodles', 'Stir-fried noodles with prawns and vegetables', 8.00, 'normal', 'chinese', 'noodle_image_url_1'),
  (2, true, 'Beef Hor Fun', 'Flat rice noodles with beef in a savory sauce', 9.50, 'normal', 'chinese', 'noodle_image_url_2'),
  (2, true, 'Char Kway Teow', 'Stir-fried flat rice noodles with prawns and chinese sausage', 10.00, 'normal', 'chinese', 'noodle_image_url_3'),
  (2, true, 'Wonton Noodle Soup', 'Noodle soup with wontons and BBQ pork', 8.50, 'normal', 'chinese', 'noodle_image_url_4'),

  -- Sushi Haven
  (3, true, 'Assorted Sushi Platter', 'Chef''s selection of fresh sushi', 12.00, 'normal', 'japanese', 'sushi_image_url_1'),
  (3, true, 'Sashimi Deluxe', 'Premium slices of assorted sashimi', 15.00, 'normal', 'japanese', 'sushi_image_url_2'),
  (3, true, 'Chirashi Bowl', 'Assorted sashimi over a bed of sushi rice', 18.00, 'normal', 'japanese', 'sushi_image_url_3'),
  (3, true, 'Dragon Roll', 'Sushi roll with eel, avocado, and cucumber', 14.00, 'normal', 'japanese', 'sushi_image_url_4'),

  -- Spicy Noodle House
  (4, true, 'Spicy Ramen', 'Fiery ramen with spicy broth', 10.00, 'normal', 'korean', 'spicy_noodle_image_url_1'),
  (4, true, 'Kimchi Udon', 'Udon noodles with kimchi and vegetables', 11.50, 'normal', 'korean', 'spicy_noodle_image_url_2'),
  (4, true, 'Bibimbap', 'Mixed rice with vegetables, beef, and a spicy sauce', 12.00, 'normal', 'korean', 'spicy_noodle_image_url_3'),
  (4, true, 'Spicy Kimchi Fried Rice', 'Fried rice with spicy kimchi and pork', 10.50, 'normal', 'korean', 'spicy_noodle_image_url_4'),
  
  -- Pizza Paradise
  (5, true, 'Margherita Pizza', 'Classic pizza with tomato, mozzarella, and basil', 14.00, 'normal', 'pizza', 'pizza_image_url_1'),
  (5, true, 'Pepperoni Delight', 'Pizza topped with pepperoni and cheese', 16.50, 'normal', 'pizza', 'pizza_image_url_2'),
  (5, true, 'Seafood Pizza', 'Pizza topped with assorted seafood and herbs', 17.00, 'normal', 'pizza', 'pizza_image_url_3'),
  (5, true, 'Vegetarian Supreme', 'Pizza loaded with assorted vegetables and cheese', 15.50, 'normal', 'pizza', 'pizza_image_url_4'),

  -- Curry House
  (6, true, 'Chicken Curry', 'Spicy chicken curry with fragrant rice', 9.00, 'normal', 'indian', 'curry_image_url_1'),
  (6, true, 'Vegetable Biryani', 'Flavorful biryani with mixed vegetables', 11.00, 'normal', 'indian', 'https://img.taste.com.au/_L7m5vxs/taste/2016/11/vegetable-biryani-102620-1.jpeg'),
  (6, true, 'Butter Chicken', 'Creamy and flavorful butter chicken curry', 12.00, 'normal', 'indian', 'https://cravinghomecooked.com/wp-content/uploads/2020/06/butter-chicken-1.jpg'),
  (6, true, 'Paneer Tikka Masala', 'Paneer in a spicy tomato-based curry', 13.50, 'normal', 'indian', 'https://healthynibblesandbits.com/wp-content/uploads/2019/07/Paneer-Tikka-Masala-5.jpg'),

  -- Thai Delights
  (7, true, 'Tom Yum Soup', 'Spicy and sour soup with shrimp', 8.50, 'normal', 'thai', 'thai_image_url_1'),
  (7, true, 'Green Curry Chicken', 'Rich green curry with tender chicken', 10.50, 'normal', 'thai', 'thai_image_url_2'),

  -- Dessert Dreamland
  (8, true, 'Mango Sticky Rice', 'Sweet sticky rice topped with fresh mango', 6.00, 'normal', 'dessert', 'dessert_image_url_1'),
  (8, true, 'Chocolate Lava Cake', 'Decadent chocolate cake with a gooey center', 7.50, 'normal', 'dessert', 'dessert_image_url_2'),
  (8, true, 'Durian Ice Cream', 'Creamy durian-flavored ice cream', 7.00, 'normal', 'dessert', 'dessert_image_url_3'),
  (8, true, 'Pandan Cake', 'Light and fluffy pandan-flavored cake', 6.50, 'normal', 'dessert', 'dessert_image_url_4'),

  -- Healthy Bites
  (9, true, 'Quinoa Salad', 'Nutrient-packed salad with quinoa and veggies', 11.00, 'normal', 'healthy', 'healthy_image_url_1'),
  (9, true, 'Grilled Salmon', 'Grilled salmon fillet with steamed broccoli', 13.50, 'normal', 'healthy', 'healthy_image_url_2'),
  (9, true, 'Avocado Salad', 'Salad with avocado, cherry tomatoes, and mixed greens', 12.50, 'normal', 'healthy', 'healthy_image_url_3'),
  (9, true, 'Teriyaki Tofu Bowl', 'Bowl with teriyaki tofu, brown rice, and vegetables', 11.00, 'normal', 'healthy', 'healthy_image_url_4'),

  -- Western Grill
  (10, true, 'BBQ Ribs', 'Tender BBQ ribs with a savory glaze', 18.00, 'normal', 'western', 'western_image_url_1'),
  (10, true, 'Classic Cheeseburger', 'Juicy beef patty with melted cheese', 14.50, 'normal', 'western', 'western_image_url_2'),
  (10, true, 'Surf and Turf', 'Steak and grilled prawns served with a side of mashed potatoes', 20.00, 'normal', 'western', 'western_image_url_3'),
  (10, true, 'Caesar Salad with Grilled Chicken', 'Classic Caesar salad with grilled chicken', 13.50, 'normal', 'western', 'western_image_url_4'),
  
  -- Satay Street
  (11, true, 'Chicken Satay', 'Grilled chicken skewers with peanut sauce', 8.00, 'normal', 'satay', 'satay_image_url_1'),
  (11, true, 'Beef Satay', 'Grilled beef skewers with a flavorful marinade', 9.00, 'normal', 'satay', 'satay_image_url_2'),
  (11, true, 'Satay Platter', 'Assortment of grilled chicken and beef satay skewers', 15.00, 'normal', 'satay', 'satay_image_url_3'),
  (11, true, 'Satay Fried Rice', 'Fried rice with chopped chicken and beef satay', 12.50, 'normal', 'satay', 'satay_image_url_4'),

  -- Dim Sum Delight
  (12, true, 'Har Gow', 'Steamed shrimp dumplings with a thin translucent skin', 5.50, 'normal', 'chinese', 'dim_sum_image_url_1'),
  (12, true, 'Siu Mai', 'Steamed pork dumplings with a savory filling', 6.50, 'normal', 'chinese', 'dim_sum_image_url_2'),
  (12, true, 'Xiao Long Bao', 'Soup dumplings with a burst of flavorful broth inside', 7.50, 'normal', 'chinese', 'dim_sum_image_url_3'),
  (12, true, 'Fried Spring Rolls', 'Crispy spring rolls filled with vegetables and shrimp', 6.00, 'normal', 'chinese', 'dim_sum_image_url_4'),

  -- Ramen House
  (13, true, 'Shoyu Ramen', 'Soy-based ramen with tender slices of pork', 10.00, 'normal', 'japanese', 'ramen_image_url_1'),
  (13, true, 'Miso Ramen', 'Ramen in a rich miso broth with vegetables', 11.50, 'normal', 'japanese', 'ramen_image_url_2'),
  (13, true, 'Spicy Tan Tan Ramen', 'Spicy sesame-based broth with ground pork', 12.00, 'normal', 'japanese', 'ramen_image_url_3'),
  (13, true, 'Vegetarian Miso Ramen', 'Ramen in miso broth with tofu and assorted vegetables', 11.00, 'normal', 'japanese', 'ramen_image_url_4');




-- INSERT INTO cuisine (cuisine)
-- VALUES;

