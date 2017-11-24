<?php

namespace PetrLatsko\JsonBundle\Client\Interfaces;

interface TransportInterface
{

    public function reset();

    public function execRequest($url);

    public function hasErrors();

    public function getErrorCode();

    public function getErrorMessage();

}
