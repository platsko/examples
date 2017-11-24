<?php

namespace PetrLatsko\JsonBundle\Client;

use PetrLatsko\JsonBundle\Client\Exceptions\TransportRequestException;

class CurlClient extends AbstractClient
{

    /**
     * @param $url
     * @throws TransportRequestException
     */
    public function execRequest($url)
    {
        $this->execTransportRequest($url);
        $this->parseJson();
    }

}
