<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;

class ListingController extends Controller
{
    public function index()
    {
        $listings = [
            ['id' => 1, 'nama' => 'Nike Court Trackzip', 'harga' => 420000, 'gender' => 'Hommes', 'size' => 'L'],
            ['id' => 2, 'nama' => 'Puma x BMW Hooded Bomber', 'harga' => 550000, 'gender' => 'Hommes', 'size' => 'M'],
            ['id' => 3, 'nama' => 'Adidas Terrex', 'harga' => 605000, 'gender' => 'Hommes', 'size' => 'XL'],
            ['id' => 4, 'nama' => 'Baselayer Rebook', 'harga' => 100000, 'gender' => 'Unisex', 'size' => 'S'],
            ['id' => 5, 'nama' => 'Kappa Trackpants', 'harga' => 75000, 'gender' => 'Femmes', 'size' => 'S'],
            ['id' => 6, 'nama' => 'Ellese Remini Tracktop', 'harga' => 650000, 'gender' => 'Unisex', 'size' => 'M'],
        ];

        return response()->json($listings);
    }
}