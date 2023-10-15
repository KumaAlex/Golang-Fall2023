ALTER TABLE laptopBags ADD CONSTRAINT laptopBags_weight_check CHECK (weight >= 0);
ALTER TABLE laptopBags ADD CONSTRAINT laptopBags_dimensions_length_check CHECK (array_length(dimensions, 1) = 3);

