INSERT INTO restaurants (name, description, address, rating, restaurant_type, num_of_reviews, image_url)
VALUES
  ('Chicken Rice Express', 'Famous for Hainanese Chicken Rice', '123 Orchard Road, Singapore', 4.9, 'Chicken', 1769, 'chicken_rice_image_url'),
  ('China Noodle Delight', 'Specializing in various noodle dishes', '456 Raffles Place, Singapore', 4.8, 'Chinese', 24322, 'noodle_image_url'),
  ('Sushi Haven', 'Authentic Japanese sushi and sashimi', '789 Marina Bay Sands, Singapore', 4.7, 'Japanese', 16381, 'sushi_image_url'),
  ('Korean Spicy Noodle House', 'Home of spicy noodle lovers', '101 Tanjong Pagar Road, Singapore', 4.6, 'Korean', 16111, 'spicy_noodle_image_url'),
  ('Pizza Paradise', 'Delicious pizzas with a variety of toppings', '202 Sentosa Island, Singapore', 4.5, 'Pizza', 7544, 'pizza_image_url'),
  ('Curry House', 'Flavorful Indian curries', '303 Little India, Singapore', 4.4, 'Indian', 16869, 'curry_image_url'),
  ('Thai Delights', 'Traditional Thai dishes with a modern twist', '404 Chinatown, Singapore', 4.3, 'Thai', 9733, 'thai_image_url'),
  ('Dessert Dreamland', 'Satisfy your sweet tooth with our desserts', '505 Bugis Street, Singapore', 4.2, 'Dessert', 4544, 'dessert_image_url'),
  ('Healthy Bites', 'Nutritious and delicious healthy options', '606 Orchard Boulevard, Singapore', 4.1, 'Healthy', 24101, 'healthy_image_url'),
  ('Western Grill', 'Classic Western dishes for all tastes', '707 Clarke Quay, Singapore', 5.0, 'Western', 19615, 'western_image_url'),
  ('Satay Street', 'Famous for delicious satay skewers', '111 Marina Bay, Singapore', 4.9, 'Satay', 21303, 'satay_image_url'),
  ('Dim Sum Delight', 'Authentic Cantonese dim sum experience', '222 Chinatown, Singapore', 4.8, 'Chinese', 2886, 'dim_sum_image_url'),
  ('Ramen House', 'Specializing in various ramen varieties', '333 Orchard Road, Singapore', 3.8, 'Japanese', 29825, 'ramen_image_url'),
  ('Mexican Maze', 'Specializing in texmex cuisines', '84 Beach Road, Singapore', 4.7, 'Japanese', 15736, 'ramen_image_url'),
  ('Vegetarian Paradise', 'Specializing in Vegetarian cuisines', '200 Toa Payoh North, Singapore', 4.6, 'Vegetarian', 15568, 'ramen_image_url'),
  ('Vegan Delights', 'Specializing in Vegan cuisines', '754 Changi Road, Singapore', 4.6, 'Vegan', 18528, 'ramen_image_url');





INSERT INTO dishes (restaurant_id, is_available, name, description, price, diet_type, cuisine, image_url)
VALUES

  -- Chicken Rice Express
  (1, true, 'Hainanese Chicken Rice', 'Classic Hainanese chicken with fragrant rice', 5.50, 'normal', '{"singaporean", "chinese", "rice"}', 'chicken_rice_image_url_1'),
  (1, true, 'Roast Chicken Noodles', 'Roasted chicken served with noodles', 6.50, 'normal', '{"singaporean", "chinese", "noodles"}', 'chicken_rice_image_url_2'),
  (1, true, 'Chicken Satay Rice', 'Hainanese-style chicken rice with chicken satay', 8.50, 'normal', '{"singaporean", "chinese", "rice"}','chicken_rice_image_url_3'),
  (1, true, 'Spicy Chicken Noodles', 'Spicy chicken noodles with a kick', 7.00, 'normal', '{"singaporean", "chinese", "noodles"}', 'chicken_rice_image_url_4'),
  (1, true, 'Chicken Rice Vermicelli Soup', 'Tender chicken, rice vermicelli noodles, and soup broth', 7.50, 'normal', '{"singaporean", "chinese", "soup", "noodles"}', 'chicken_rice_image_url_5'),
  (1, true, 'Black Pepper Chicken Rice Bowl', 'Savory chicken in a black pepper sauce served over rice', 9.00, 'normal', '{"singaporean", "chinese", "rice"}', 'chicken_rice_image_url_6'),
  (1, true, 'Sesame Ginger Chicken Salad', 'Grilled chicken with sesame ginger dressing on a bed of greens', 8.00, 'normal', '{"singaporean", "chinese", "healthy"}', 'chicken_rice_image_url_7'),
  (1, true, 'Chicken Rice Paper Rolls', 'Fresh rice paper rolls filled with chicken, herbs, and vegetables', 6.50, 'normal', '{"singaporean", "chinese"}', 'chicken_rice_image_url_8'),
  (1, true, 'Curry Chicken Rice', 'Tender chicken in a fragrant curry sauce served with rice', 8.50, 'normal', '{"singaporean", "chinese"}', 'chicken_rice_image_url_9'),
  (1, true, 'Lemon Herb Roast Chicken', 'Juicy roast chicken seasoned with lemon and herbs', 10.00, 'normal', '{"singaporean", "chinese", "roast"}', 'chicken_rice_image_url_10'),
  (1, true, 'Chicken Congee', 'Traditional rice porridge with shredded chicken and aromatics', 6.00, 'normal', '{"singaporean", "chinese", "rice"}', 'chicken_rice_image_url_11'),
  (1, true, 'Ginger Scallion Chicken Rice', 'Steamed chicken with ginger scallion sauce over rice', 7.50, 'normal', '{"singaporean", "chinese", "rice"}', 'chicken_rice_image_url_12'),

  -- China Noodle Delight
  (2, true, 'Singapore-style Noodles', 'Stir-fried noodles with prawns and vegetables', 8.00, 'normal', '{"singaporean", "chinese", "noodles"}', 'noodle_image_url_1'),
  (2, true, 'Beef Hor Fun', 'Flat rice noodles with beef in a savory sauce', 9.50, 'normal', '{"singaporean", "chinese", "noodles"}', 'noodle_image_url_2'),
  (2, true, 'Char Kway Teow', 'Stir-fried flat rice noodles with prawns and chinese sausage', 10.00, 'normal', '{"singaporean", "chinese", "noodles"}', 'noodle_image_url_3'),
  (2, true, 'Wonton Noodle Soup', 'Noodle soup with wontons and BBQ pork', 8.50, 'normal', '{"singaporean", "chinese", "noodles"}', 'noodle_image_url_4'),
  (2, true, 'Dan Dan Noodles', 'Spicy Sichuan noodles with minced pork and peanuts', 9.00, 'normal', '{"chinese", "noodles"}', 'noodle_image_url_5'),
  (2, true, 'Vegetable Lo Mein', 'Stir-fried noodles with assorted vegetables in a light soy sauce', 7.50, '{"vegetarian", "normal"}', '{"chinese", "noodles", "singaporean"}', 'noodle_image_url_6'),
  (2, true, 'Spicy Garlic Eggplant Noodles', 'Eggplant and noodles in a spicy garlic sauce', 8.50, 'normal', '{"chinese", "noodles", "spicy", "singaporean"}', 'noodle_image_url_7'),
  (2, true, 'Cold Sesame Noodles', 'Chilled noodles tossed in a sesame sauce with cucumber and sesame seeds', 8.00, 'normal', '{"chinese", "noodles", "cold"}', 'noodle_image_url_8'),
  (2, true, 'Shanghai Stir-Fried Udon', 'Thick udon noodles stir-fried with pork, cabbage, and soy sauce', 9.50, 'normal', '{"chinese", "noodles", "udon"}', 'noodle_image_url_9'),
  (2, true, 'Seafood Noodle Soup', 'Noodle soup with assorted seafood and fish balls in a clear broth', 10.50, 'normal', '{"chinese", "noodles", "seafood", "singaporean"}', 'noodle_image_url_10'),
  (2, true, 'Szechuan Spicy Noodles', 'Hand-pulled noodles in a fiery Szechuan pepper sauce with minced meat', 9.00, 'normal', '{"chinese", "noodles", "spicy"}', 'noodle_image_url_11'),
  (2, true, 'Hong Kong Soy Sauce Noodles', 'Egg noodles tossed in a savory soy sauce with char siu pork', 9.00, 'normal', '{"chinese", "noodles"}', 'noodle_image_url_12'),
  (2, true, 'Vegetable Chow Mein', 'Stir-fried noodles with mixed vegetables in a savory sauce', 7.00, '{"vegetarian", "normal"}', '{"chinese", "noodles", "singaporean"}', 'vegetarian_noodle_image_url_1'),
  (2, true, 'Tofu and Veggie Stir-Fry Noodles', 'Tofu and assorted veggies stir-fried with noodles in a light sauce', 8.50, '{"vegetarian", "normal"}', '{"chinese", "noodles", "singaporean"}', 'vegetarian_noodle_image_url_2'),


  -- Sushi Haven
  (3, true, 'Assorted Sushi Platter', 'Chef''s selection of fresh sushi', 12.00, 'normal', '{"japanese", "rice"}', 'sushi_image_url_1'),
  (3, true, 'Sashimi Deluxe', 'Premium slices of assorted sashimi', 15.00, 'normal', 'japanese', 'sushi_image_url_2'),
  (3, true, 'Chirashi Bowl', 'Assorted sashimi over a bed of sushi rice', 18.00, 'normal', '{"japanese", "rice"}', 'sushi_image_url_3'),
  (3, true, 'Dragon Roll', 'Sushi roll with eel, avocado, and cucumber', 14.00, 'normal', '{"japanese", "rice"}', 'sushi_image_url_4'),
  (3, true, 'Vegetable Tempura Roll', 'Crispy tempura-battered vegetables in a sushi roll', 10.00, '{"vegetarian", "normal"}', '{"japanese", "rice"}', 'vegetarian_sushi_image_url_5'),
  (3, true, 'Avocado Cucumber Roll', 'Sushi roll filled with creamy avocado and crisp cucumber', 9.00, '{"vegetarian", "normal"}', '{"japanese", "rice"}', 'vegetarian_sushi_image_url_6'),
  (3, true, 'Inari Sushi', 'Sweet tofu pockets filled with seasoned sushi rice', 8.50, '{"vegetarian", "normal"}', '{"japanese", "rice"}', 'vegetarian_sushi_image_url_7'),
  (3, true, 'Miso Glazed Eggplant Nigiri', 'Grilled eggplant nigiri with a savory miso glaze', 9.50, '{"vegetarian", "normal"}', '{"japanese", "rice"}', 'vegetarian_sushi_image_url_8'),
  (3, true, 'Vegetable Hand Roll', 'Hand roll with fresh vegetables and sushi rice wrapped in seaweed', 7.50, '{"vegetarian", "normal"}', '{"japanese", "rice"}', 'vegetarian_sushi_image_url_9'),
  (3, true, 'Crispy Tofu Maki', 'Crunchy tempura tofu maki roll with tangy dipping sauce', 11.00, '{"vegetarian", "normal"}', '{"japanese", "rice"}', 'vegetarian_sushi_image_url_10'),
  (3, true, 'Asparagus Nigiri', 'Fresh asparagus atop seasoned sushi rice', 8.00, '{"vegetarian", "normal"}', '{"japanese", "rice"}', 'vegetarian_sushi_image_url_11'),
  (3, true, 'Pickled Radish Roll', 'Colorful roll with pickled radish, avocado, and sushi rice', 9.00, '{"vegetarian", "normal"}', '{"japanese", "rice"}', 'vegetarian_sushi_image_url_12'),
  (3, true, 'Salmon Nigiri', 'Fresh salmon slices over seasoned sushi rice', 9.00, 'normal', '{"japanese", "rice"}', 'non_vegetarian_sushi_image_url_13'),
  (3, true, 'Tuna Sashimi', 'Buttery tuna slices served sashimi-style', 12.00, 'normal', '{"japanese", "sashimi", "tuna"}', 'non_vegetarian_sushi_image_url_14'),
  (3, true, 'Spicy Tuna Roll', 'Sushi roll with spicy tuna and avocado', 11.00, 'normal', '{"japanese", "rice"}', 'non_vegetarian_sushi_image_url_15'),
  (3, true, 'Prawn Tempura Uramaki', 'Inside-out sushi roll with crispy prawn tempura', 13.00, 'normal', '{"japanese", "rice"}', 'non_vegetarian_sushi_image_url_16'),



  -- Korean Spicy Noodle House
  (4, true, 'Spicy Ramen', 'Fiery ramen with spicy broth', 10.00, 'normal', '{"korean", "noodles"}', 'spicy_noodle_image_url_1'),
  (4, true, 'Kimchi Udon', 'Udon noodles with kimchi and vegetables', 11.50, 'normal', '{"korean", "noodles"}', 'spicy_noodle_image_url_2'),
  (4, true, 'Bibimbap', 'Mixed rice with vegetables, beef, and a spicy sauce', 12.00, 'normal', '{"korean", "rice"}', 'spicy_noodle_image_url_3'),
  (4, true, 'Spicy Kimchi Fried Rice', 'Fried rice with spicy kimchi and pork', 10.50, 'normal', '{"korean", "rice"}', 'spicy_noodle_image_url_4'),
  
  -- Pizza Paradise
  (5, true, 'Margherita Pizza', 'Classic pizza with tomato, mozzarella, and basil', 14.00, 'normal', '{"italian", "american"}', 'pizza_image_url_1'),
  (5, true, 'Pepperoni Delight', 'Pizza topped with pepperoni and cheese', 16.50, 'normal', '{"italian", "american"}', 'pizza_image_url_2'),
  (5, true, 'Seafood Pizza', 'Pizza topped with assorted seafood and herbs', 17.00, 'normal', '{"italian", "american"}', 'pizza_image_url_3'),
  (5, true, 'Vegetarian Supreme', 'Pizza loaded with assorted vegetables and cheese', 15.50, '{"vegetarian", "normal"}', '{"italian", "american"}', 'pizza_image_url_4'),

  -- Curry House
  (6, true, 'Chicken Curry', 'Spicy chicken curry with fragrant rice', 9.00, '{"halal", "normal"}', '{"indian", "rice"}', 'https://www.recipetineats.com/wp-content/uploads/2023/10/African-coconut-chicken-curry_3.jpg?w=500&h=500&crop=1'),
  (6, true, 'Vegetable Biryani', 'Flavorful biryani with mixed vegetables', 11.00, '{"halal", "vegetatian", "normal"}', '{"indian", "rice"}', 'https://img.taste.com.au/_L7m5vxs/taste/2016/11/vegetable-biryani-102620-1.jpeg'),
  (6, true, 'Butter Chicken', 'Creamy and flavorful butter chicken curry', 12.00, '{"halal", "normal"}', 'indian', 'https://cravinghomecooked.com/wp-content/uploads/2020/06/butter-chicken-1.jpg'),
  (6, true, 'Paneer Tikka Masala', 'Paneer in a spicy tomato-based curry', 13.50, '{"halal", "vegetarian", "normal"}', '{"indian", "rice"}', 'https://healthynibblesandbits.com/wp-content/uploads/2019/07/Paneer-Tikka-Masala-5.jpg'),
   (6, true, 'Masala Dosa', 'Crispy crepe filled with spiced potatoes', 8.50, '{"halal", "vegetarian", "normal"}', '{"indian"}', 'https://recipes.net/wp-content/uploads/2023/05/masala-dosa-recipe_d563150be0b0109e2b9ef6e2d6a51c0e.jpeg'),
  (6, true, 'Chana Masala', 'Spiced chickpea curry with aromatic spices', 10.00, '{"halal", "vegetarian", "normal"}', '{"indian"}', 'https://www.indianhealthyrecipes.com/wp-content/uploads/2021/08/chana-masala-recipe.jpg'),
  (6, true, 'Palak Paneer', 'Cottage cheese cubes in a creamy spinach curry', 12.00, '{"halal", "vegetarian", "normal"}', '{"indian"}', 'https://healthynibblesandbits.com/wp-content/uploads/2020/01/Saag-Paneer-FF.jpg'),
  (6, true, 'Aloo Gobi', 'Potato and cauliflower cooked with Indian spices', 9.00, '{"halal", "vegetarian", "normal"}', '{"indian"}', 'https://images.immediate.co.uk/production/volatile/sites/30/2018/04/Roasted_aloo_gobi-a6671e9-scaled.jpg?quality=90&resize=556,505'),
  (6, true, 'Vegetable Korma', 'Assorted vegetables in a creamy, nutty sauce', 11.50, '{"halal", "vegetarian", "normal"}', '{"indian"}', 'https://myfancypantry.files.wordpress.com/2012/04/untitled-04473.jpg'),
  (6, true, 'Chicken Biryani', 'Flavorful rice cooked with marinated chicken', 14.00, '{"halal", "normal"}', '{"indian", "rice"}', 'https://www.awesomecuisine.com/wp-content/uploads/2007/10/Chicken-Biryani_resized-500x436.jpg'),
  (6, true, 'Mutton Rogan Josh', 'Slow-cooked mutton in a rich, flavorful gravy', 15.00, '{"halal", "normal"}', 'indian', 'https://www.archanaskitchen.com/images/archanaskitchen/1-Author/Shaheen_Ali/Kashmiri_Rogan_Josh.jpg'),
  (6, true, 'Pani Puri', 'Indian street food snack with spiced water and chickpeas', 6.50, '{"halal", "vegetarian", "normal"}', '{"indian"}', 'https://www.sidechef.com/recipe/3883dffb-5fa2-4ee9-8054-d8de1409899f.jpg'),
  (6, true, 'Rasgulla', 'Spongy Bengali sweet dumplings in sugar syrup', 5.00, '{"halal", "vegetarian", "normal"}', '{"indian", "dessert"}', 'https://cdn.ready-market.com.tw/21cd62de/Templates/pic/ANKO-Rasgulla-1200X1200.jpg?v=1a046a18'),
  (6, true, 'Kadai Paneer', 'Paneer and bell peppers in a spicy tomato-based gravy', 13.00, '{"halal", "vegetarian", "normal"}', '{"indian"}', 'https://www.secondrecipe.com/wp-content/uploads/2020/05/dhaba-style-kadai-paneer-001.jpg'),

  -- Thai Delights
  (7, true, 'Tom Yum Soup', 'Spicy and sour soup with shrimp', 8.50, 'normal', '{"thai", "soup"}', 'https://hot-thai-kitchen.com/wp-content/uploads/2013/03/tom-yum-goong-blog.jpg'),
  (7, true, 'Green Curry Chicken', 'Rich green curry with tender chicken', 10.50, 'normal', '{"thai", "soup"}', 'https://hot-thai-kitchen.com/wp-content/uploads/2022/04/green-curry-new-sq-3.jpg'),
  (7, true, 'Vegetarian Pad Thai', 'Stir-fried rice noodles with tofu and vegetables', 9.00, '{"vegetarian", "normal"}', '{"thai", "noodles"}', 'https://img.taste.com.au/z9EIVHJg/taste/2021/02/10-minute-vegetarian-pad-thai-168946-2.jpg'),
  (7, true, 'Vegan Phat Kaphrao', 'Spicy stir-fried tofu with holy basil', 11.00, '{"vegan", "vegetarian", "normal"}', '{"thai", "stir-fry"}', 'https://www.alphafoodie.com/wp-content/uploads/2020/09/vegan-thai-basil-beef-1-of-1-3.jpeg'),
  (7, true, 'Som Tum Salad', 'Refreshing green papaya salad with peanuts and chili', 8.50, '{"vegetarian", "normal"}', '{"thai", "healthy"}', 'https://www.seriouseats.com/thmb/yKNZ9ICJC5ZNhzcYHdHENxogpFw=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc()/20210928-Som-Tam-Thai-green-papaya-salad-vicky-wasik-24-f0d666fc609f49a0b9f34897bd2c6303.jpg'),
  (7, true, 'Vegan Red Curry', 'Creamy red curry with mixed vegetables and tofu', 12.50, '{"vegan", "vegetarian", "normal"}', '{"thai", "curry"}', 'https://www.connoisseurusveg.com/wp-content/uploads/2023/04/vegan-red-curry-sq-1.jpg'),
  (7, true, 'Vegetable Spring Rolls', 'Crispy fried spring rolls filled with veggies', 6.50, '{ "vegetarian", "normal"}', '{"thai", "appetizer"}', 'https://www.thecookingcollective.com.au/wp-content/uploads/2022/08/Vegetable-Spring-Rolls-2-1.jpg'),
  (7, true, 'Vegan Larb', 'Spicy minced tofu salad with herbs and lime dressing', 10.00, '{"vegan", "vegetarian", "normal"}', '{"thai", "salad"}', 'https://static01.nyt.com/images/2021/08/11/dining/11Veggie-tofu-larb-a/merlin_188546805_b6a2a32b-5635-46e5-82b6-b11b7ed5ac3f-articleLarge.jpg'),
  (7, true, 'Vegetarian Pad See Ew', 'Stir-fried rice noodles with broccoli and egg', 9.50, '{"vegetarian", "normal"}', '{"thai", "noodles"}', 'https://www.myplantifulcooking.com/wp-content/uploads/2022/11/vegan-pad-see-ew-featured.jpg'),
  (7, true, 'Vegan Massaman Curry', 'Rich and nutty curry with potatoes and tofu', 13.00, '{"vegan", "vegetarian", "normal"}', '{"thai", "curry"}', 'https://images.kitchenstories.io/wagtailOriginalImages/R2456-new-final-photo-2.jpg'),
  (7, true, 'Green Papaya Salad', 'Zesty salad made with shredded green papaya, tomatoes, and peanuts', 7.50, '{"vegetarian", "normal"}', '{"thai", "salad"}', 'https://www.whiskaffair.com/wp-content/uploads/2020/06/Green-Papaya-Salad-2-3.jpg'),
  (7, true, 'Vegan Pad Kaprow', 'Spicy stir-fried tofu with Thai basil', 11.50, '{"vegan", "vegetarian", "normal"}', '{"thai", "rice"}', 'https://joyfuldumplings.com/wp-content/uploads/2021/07/Vegan_Thai_Basil_Stir-fry_Pad_Kra_Pao_3_11zon.jpg'),
  (7, true, 'Beef Panang Curry', 'Tender beef in a rich and spicy Panang curry sauce', 14.50, 'normal', '{"thai", "curry"}', 'https://www.foodandwine.com/thmb/atro54_2GCQAuL_BfcWOHg094xs=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc()/Penang-Curry-FT-RECIPE0523-4bd3376f76a648b3b8a30929df61acc5.jpg'),
  (7, true, 'Chicken Satay', 'Grilled marinated chicken skewers served with peanut sauce', 10.00, 'normal', '{"thai", "appetizer"}', 'https://food.fnr.sndimg.com/content/dam/images/food/plus/fullset/2020/02/06/0/FNK-Live_112119-Jet-Tila-Chicken-Satay-with-Peanut-Sauce-0005_s4x3.jpg.rend.hgtvcom.616.462.suffix/1581032331905.jpeg'),
  (7, true, 'Pork Basil Stir Fry', 'Stir-fried pork with Thai basil and spicy sauce', 12.00, 'normal', '{"thai", "stir-fry"}', 'https://images.services.kitchenstories.io/MRxcM_KNscGAJaphGq57qFKQWJ8=/3840x0/filters:quality(85)/images.kitchenstories.io/wagtailOriginalImages/R2592-final-photo.jpg'),
  (7, true, 'Thai Iced Tea', 'Sweet and creamy iced tea made with condensed milk', 3.50, '{"vegan", "vegetarian", "normal"}', '{"thai", "drink"}', 'https://www.thespruceeats.com/thmb/w16B4MH-a_AlvG4t27cw3KgrjLM=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc()/SES-real-thai-iced-tea-recipe-765458-hero-1-9139caef4772411bb34a948c5df50cfb.jpg'),
  (7, true, 'Fresh Coconut Water', 'Refreshing coconut water served straight from the coconut', 4.00, '{"vegan", "vegetarian", "normal"}', '{"thai", "drink"}', 'https://images.squarespace-cdn.com/content/v1/5c1074accc8fed6a4251da8f/4882b78a-b5d4-47a1-a9a0-d05a4fdbd724/shutterstock_490174816.jpg'),
  (7, true, 'Mango Sticky Rice Smoothie', 'Indulgent smoothie inspired by the popular Thai dessert', 5.50, '{"vegan", "vegetarian", "normal"}', '{"thai", "drink"}', 'https://media-cdn.tripadvisor.com/media/photo-p/12/87/06/dc/mango-sticky-rice-smoothie.jpg');
  (7, true, 'Thai Lemongrass Tea', 'Fragrant tea made with lemongrass and spices', 3.00, '{"vegan", "vegetarian", "normal"}', '{"thai", "drink"}', 'https://www.tastythais.com/wp-content/uploads/2020/02/thai-lemongrass-tea.jpg'),
  (7, true, 'Watermelon Cooler', 'Refreshing watermelon drink with a hint of mint', 4.50, '{"vegan", "vegetarian", "normal"}', '{"thai", "drink"}', 'https://anothertablespoon.com/wp-content/uploads/2018/06/DSC01615-scaled.jpg'),
  (7, true, 'Thai Basil Seed Drink', 'Healthy and soothing drink made with basil seeds', 3.50, '{"vegan", "vegetarian", "normal"}', '{"thai", "drink"}', 'https://foodnaturalthai.igetweb.com/catalog/p_1473291.jpg'),

  -- Dessert Dreamland
  (8, true, 'Mango Sticky Rice', 'Sweet sticky rice topped with fresh mango', 6.00, '{"halal", "vegan", "vegetarian", "normal"}', '{"dessert", "rice"}', 'dessert_image_url_1'),
  (8, true, 'Chocolate Lava Cake', 'Decadent chocolate cake with a gooey center', 7.50, '{"halal", "vegan", "vegetarian", "normal"}', 'dessert', 'dessert_image_url_2'),
  (8, true, 'Durian Ice Cream', 'Creamy durian-flavored ice cream', 7.00, '{"halal", "vegan", "vegetarian", "normal"}', 'dessert', 'dessert_image_url_3'),
  (8, true, 'Pandan Cake', 'Light and fluffy pandan-flavored cake', 6.50, '{"halal", "vegan", "vegetarian", "normal"}', 'dessert', 'dessert_image_url_4'),

  -- Healthy Bites
  (9, true, 'Quinoa Salad', 'Nutrient-packed salad with quinoa and veggies', 11.00, '{"halal", "vegan", "vegetarian", "normal"}', 'healthy', 'healthy_image_url_1'),
  (9, true, 'Grilled Salmon', 'Grilled salmon fillet with steamed broccoli', 13.50, '{"halal", "vegan", "vegetarian", "normal"}', 'healthy', 'healthy_image_url_2'),
  (9, true, 'Avocado Salad', 'Salad with avocado, cherry tomatoes, and mixed greens', 12.50, '{"halal", "vegan", "vegetarian", "normal"}', 'healthy', 'healthy_image_url_3'),
  (9, true, 'Teriyaki Tofu Bowl', 'Bowl with teriyaki tofu, brown rice, and vegetables', 11.00, '{"halal", "vegan", "vegetarian", "normal"}', 'healthy', 'healthy_image_url_4'),

  -- Western Grill
  (10, true, 'BBQ Ribs', 'Tender BBQ ribs with a savory glaze', 18.00, 'normal', '{"western", "american"}', 'western_image_url_1'),
  (10, true, 'Classic Cheeseburger', 'Juicy beef patty with melted cheese', 14.50, 'normal', '{"western", "american"}', 'western_image_url_2'),
  (10, true, 'Surf and Turf', 'Steak and grilled prawns served with a side of mashed potatoes', 20.00, '{"western", "american"}', 'western', 'western_image_url_3'),
  (10, true, 'Caesar Salad with Grilled Chicken', 'Classic Caesar salad with grilled chicken', 13.50, '{"western", "american", "healthy"}', 'western', 'western_image_url_4'),
  
  -- Satay Street
  (11, true, 'Chicken Satay', 'Grilled chicken skewers with peanut sauce', 8.00, '{"halal", "normal"}', '{"indonesian", "singaporean"}', 'satay_image_url_1'),
  (11, true, 'Beef Satay', 'Grilled beef skewers with a flavorful marinade', 9.00, '{"halal", "normal"}', '{"indonesian", "singaporean"}', 'satay_image_url_2'),
  (11, true, 'Satay Platter', 'Assortment of grilled chicken and beef satay skewers', 15.00, '{"halal", "normal"}', '{"indonesian", "singaporean"}', 'satay_image_url_3'),
  (11, true, 'Satay Fried Rice', 'Fried rice with chopped chicken and beef satay', 12.50, '{"halal", "normal"}', '{"indonesian", "singaporean", "rice"}', 'satay_image_url_4'),

  -- Dim Sum Delight
  (12, true, 'Har Gow', 'Steamed shrimp dumplings with a thin translucent skin', 5.50, 'normal', '{"chinese", "singaporean"}', 'dim_sum_image_url_1'),
  (12, true, 'Siu Mai', 'Steamed pork dumplings with a savory filling', 6.50, 'normal', '{"chinese", "singaporean"}', 'dim_sum_image_url_2'),
  (12, true, 'Xiao Long Bao', 'Soup dumplings with a burst of flavorful broth inside', 7.50, 'normal', '{"chinese", "singaporean"}', 'dim_sum_image_url_3'),
  (12, true, 'Fried Spring Rolls', 'Crispy spring rolls filled with vegetables and shrimp', 6.00, 'normal', '{"chinese", "singaporean"}', 'dim_sum_image_url_4'),
  (12, true, 'Vegetarian Spring Rolls', 'Crispy spring rolls filled with vegetables', 6.00, 'vegetarian', '{"chinese", "singaporean"}', 'dim_sum_image_url_4'),

  -- Ramen House
  (13, true, 'Shoyu Ramen', 'Soy-based ramen with tender slices of pork', 10.00, 'normal', '{"japanese", "noodles"}', 'ramen_image_url_1'),
  (13, true, 'Miso Ramen', 'Ramen in a rich miso broth with vegetables', 11.50, 'normal', '{"japanese", "noodles"}', 'ramen_image_url_2'),
  (13, true, 'Spicy Tan Tan Ramen', 'Spicy sesame-based broth with ground pork', 12.00, 'normal', '{"japanese", "noodles"}', 'ramen_image_url_3'),
  (13, true, 'Vegetarian Miso Ramen', 'Ramen in miso broth with tofu and assorted vegetables', 11.00, 'vegetarian', '{"japanese", "noodles", "healthy"}', 'ramen_image_url_4'),

  -- Mexican Maze
  (14, true, 'Beef Fajitas', 'Sizzling beef strips with peppers and onions, served with tortillas', 14.00, 'normal', '{"mexican", "western"}', 'beef_fajitas_image_url_1'),
  (14, true, 'Chicken Enchiladas', 'Corn tortillas filled with seasoned chicken, cheese, and topped with sauce', 12.50, 'normal',  '{"mexican", "western"}', 'chicken_enchiladas_image_url_2'),
  (14, true, 'Queso Dip', 'Creamy cheese dip served with tortilla chips', 8.50, '{"vegetarian", "normal"}',  '{"mexican", "western"}', 'queso_dip_image_url_3'),
  (14, true, 'Chili Con Carne', 'Hearty beef and bean chili with a hint of spice', 10.00, 'normal',  '{"mexican", "western"}', 'chili_con_carne_image_url_4'),
  (14, true, 'Fish Tacos', 'Crispy fish in soft tortillas with slaw and creamy dressing', 13.00, 'normal',  '{"mexican", "western"}', 'fish_tacos_image_url_5'),
  (14, true, 'Vegetarian Burrito Bowl', 'Rice, beans, vegetables, and toppings in a bowl', 11.00, '{"vegan", "vegetarian", "normal"}', '{"mexican", "western"}', 'vegetarian_burrito_bowl_image_url_6'),
  (14, true, 'Tamale Pie', 'Cornmeal crust filled with seasoned meat, beans, and cheese', 9.50, 'normal',  '{"mexican", "western"}', 'tamale_pie_image_url_7'),
  (14, true, 'Mexican Street Corn', 'Grilled corn topped with cheese, spices, and herbs', 6.50, 'vegetarian',  '{"mexican", "western"}', 'mexican_street_corn_image_url_8'),
  (14, true, 'Sopapillas', 'Fried pastry topped with cinnamon and honey', 7.00, '{"vegan", "vegetarian", "normal"}',  '{"mexican", "western"}', 'sopapillas_image_url_9'),
  (14, true, 'Margarita Cocktail', 'Classic tequila-based cocktail with lime and salt rim', 8.00, 'normal',  '{"mexican", "western", "drink"}', 'margarita_cocktail_image_url_10'),

 -- Vegetarian Paradise
  (15, true, 'Vegetable Stir-Fry', 'Assorted fresh vegetables stir-fried in a savory sauce', 12.00, '{"vegan", "vegetarian", "normal"}', '{"chinese", "italian", "healthy"}', 'vegetable_stir_fry_image_url_1'),
  (15, true, 'Mushroom Risotto', 'Creamy Italian risotto cooked with mushrooms and parmesan cheese', 14.50, '{"vegetarian", "normal"}', '{"western", "italian", "healthy"}', 'mushroom_risotto_image_url_2'),
  (15, true, 'Spinach and Ricotta Stuffed Shells', 'Pasta shells filled with spinach, ricotta, and topped with marinara sauce', 13.00, '{"vegan", "vegetarian", "normal"}', '{"western", "italian", "mexican", "healthy"}', 'stuffed_shells_image_url_3'),
  (15, true, 'Falafel Platter', 'Crispy chickpea patties served with hummus, pita, and salad', 11.00, '{"vegetarian", "normal"}', '{"western", "indian", "healthy"}', 'falafel_platter_image_url_4'),
  (15, true, 'Eggplant Parmesan', 'Layered eggplant slices baked with marinara sauce and melted cheese', 12.00, '{"vegetarian", "normal"}', '{"western", "italian", "chinese", "healthy"}', 'eggplant_parmesan_image_url_5'),
  (15, true, 'Quinoa Salad', 'Colorful salad with quinoa, mixed vegetables, feta cheese, and a citrus dressing', 10.50, '{"vegan", "vegetarian", "normal"}', '{"western", "italian", "healthy"}', 'quinoa_salad_image_url_6'),
  (15, true, 'Vegetarian Pad Thai', 'Thai stir-fried noodles with tofu, veggies, and a tangy sauce', 13.50, '{"vegan", "vegetarian", "normal"}', '{"western", "thai", "healthy"}', 'vegetarian_pad_thai_image_url_15'),
  (15, true, 'Portobello Mushroom Burger', 'Juicy grilled portobello mushroom served on a bun with toppings', 11.50, '{"vegan", "vegetarian", "normal"}', '{"western", "american", "healthy"}', 'mushroom_burger_image_url_8'),
  (15, true, 'Coconut Curry Lentils', 'Creamy coconut curry with lentils, served with rice', 11.50, '{"vegetarian", "normal"}', '{"indian", "healthy"}', 'coconut_curry_lentils_image_url_9'),
  (15, true, 'Avocado Toast', 'Crusty bread topped with mashed avocado, cherry tomatoes, and a sprinkle of seasoning', 9.00, '{"vegan", "vegetarian", "normal"}', '{"western", "italian", "healthy"}', 'avocado_toast_image_url_10'),


-- Vegan Delights
(16, true, 'Spicy Cauliflower Wings', 'Crispy baked cauliflower florets tossed in spicy sauce', 10.00, '{"vegan", "vegetarian", "normal"}', '{"vegan", "chinese", "healthy", "singaporean"}', 'cauliflower_wings_image_url_1'),
(16, true, 'Lentil Shepherds Pie', 'Hearty lentil filling topped with mashed potatoes, baked to perfection', 14.50, '{"vegan", "vegetarian", "normal"}', '{"vegan", "western", "singaporean"}', 'shepherds_pie_image_url_2'),
(16, true, 'Vegan Tacos', 'Black bean, corn, and avocado-filled tacos served with salsa', 12.00, '{"vegan", "vegetarian", "normal"}', '{"vegan", "itatian", "western", "healthy"}', 'vegan_tacos_image_url_3'),
(16, true, 'Chickpea Shawarma Wrap', 'Marinated chickpeas with tahini sauce wrapped in a warm pita', 11.00, '{"vegan", "vegetarian", "normal"}', '{"vegan", "indian", "western", "healthy"}', 'shawarma_wrap_image_url_4'),
(16, true, 'Quinoa Stuffed Bell Peppers', 'Bell peppers filled with quinoa, black beans, and spices', 13.00, '{"vegan", "vegetarian", "normal"}', '{"vegan", "indian", "western", "singaporean"}', 'stuffed_bell_peppers_image_url_5'),
(16, true, 'Coconut Curry Tofu', 'Creamy coconut curry with crispy tofu, served with rice', 12.00, '{"vegan", "vegetarian", "normal"}', '{"vegan", "singaporean", "chinese", "healthy"}', 'coconut_curry_tofu_image_url_6'),
(16, true, 'Vegan Pizza', 'Thin-crust pizza topped with veggies, dairy-free cheese, and marinara sauce', 13.50, '{"vegan", "vegetarian", "normal"}', '{"vegan", "italian", "western"}', 'vegan_pizza_image_url_7'),
(16, true, 'Zucchini Noodles with Pesto', 'Spiralized zucchini noodles tossed in homemade vegan pesto sauce', 11.50, '{"vegan", "vegetarian", "normal"}', '{"vegan", "italian", "western", "noodles"}', 'zucchini_noodles_pesto_image_url_16'),
(16, true, 'Vegan Sushi Roll', 'Colorful sushi roll filled with avocado, cucumber, and pickled veggies', 12.50, '{"vegan", "vegetarian", "normal"}', '{"vegan", "healthy", "japanese"}', 'vegan_sushi_roll_image_url_9'),
(16, true, 'Chocolate Avocado Mousse', 'Decadent chocolate mousse made with avocado for a creamy texture', 8.00, '{"vegan", "vegetarian", "normal"}', '{"vegan", "dessert", "western"}', 'chocolate_mousse_image_url_10');



-- INSERT INTO cuisine (cuisine)
-- VALUES;

